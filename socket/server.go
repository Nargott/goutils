package socket

import (
	"net"
	"fmt"
	"os"
)

type ServerState int

type Server struct {
	name          string
	listener      net.Listener
	connectionPtr *net.Conn
	error         chan error
	data          chan string
	state         ServerState
}

const (
	//server socket states
	ServerStateInit               ServerState = iota
	ServerStateListen
	ServerStateClientConnected
	ServerStateClientDisconnected
	ServerStateStopped
)

func NewSocketServer(socketName string, dataChan chan string, errorChan chan error) *Server {
	return &Server{
		name:  fmt.Sprintf(nameMask, socketName),
		data:  dataChan,
		error: errorChan,
		state: ServerStateInit,
	}
}

func (this *Server) Listen() error {
	l, err := net.Listen(network, this.name)
	if err != nil {
		return fmt.Errorf("socket listen error: %s", err)
	}

	this.listener = l
	this.state = ServerStateListen

	return nil
}

func (this *Server) readInput() {
	//wait for connection
	fd, err := this.listener.Accept()
	if err != nil {
		this.error <- fmt.Errorf("accept error: %s", err.Error())
	}

	*this.connectionPtr = fd
	this.state = ServerStateClientConnected

	buf := make([]byte, ReadBufSize)
	for { //process input data
		if this.connectionPtr != nil {
			fd := *this.connectionPtr
			nr, err := fd.Read(buf)
			if err != nil {
				this.error <- fmt.Errorf("read data error: %s", err.Error())
			}
			this.data <- string(buf[0:nr])
		} else {
			//TODO set StateClientDisconnected somehow another
			this.state = ServerStateClientDisconnected
		}
	}
}

func (this *Server) Open() error {
	err := this.Listen()
	if err != nil {
		return err
	}
	if this.state == ServerStateListen {
		go this.readInput()
	}
}

func (this *Server) Close() {
	if this.state != ServerStateStopped {
		this.SendStopCmd()
		if this.connectionPtr != nil {
			fd := *this.connectionPtr
			fd.Close()
		}
		this.listener.Close()
		os.Remove(this.name)
		this.state = ServerStateStopped
	}
}

func (this *Server) SendData(cmd string, data interface{}) {
	if this.state == ServerStateClientConnected && this.connectionPtr != nil {
		con := *this.connectionPtr
		packet, err := packSocketDataPacket(cmd, data)
		if err != nil {
			this.error <- err
		}
		_, err = con.Write(packet)
		if err != nil {
			this.error <- fmt.Errorf("write ('stop' command) error: %s", err)
		}
	}
}

func (this *Server) SendStopCmd() {
	this.SendData(CmdStop)
}

func (this *Server) SendOkCmd() {
	this.SendData(CmdOk)
}

func (this *Server) SendErrorCmd(msg string) {
	this.SendData(CmdError, msg) //TODO add message may be
}

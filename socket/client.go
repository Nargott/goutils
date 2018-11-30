package socket

import (
	"net"
	"fmt"
)

type ClientState int

type Client struct {
	name          string
	listener      net.Listener
	connectionPtr *net.Conn
	error         chan error
	data          chan string
	state         ClientState
}

const (
	//client socket states
	ClientStateInit               ClientState = iota
	ClientStateConnected
	ClientStateStopped
)

func NewSocketClient(socketName string) *Client {
	return &Client{
		name:  fmt.Sprintf(nameMask, socketName),
		state: ClientStateInit,
	}
}

func (this *Client) connect() error {
	if this.state == ClientStateInit || this.state == ClientStateStopped {
		c, err := net.Dial(network, this.name)
		if err != nil {
			return err
		}

		this.connectionPtr = &c
		this.state = ClientStateConnected
	}

	return nil
}

func (this *Client) readInput() error {
	if this.state == ClientStateConnected && this.connectionPtr != nil {
		buf := make([]byte, 1024)
		c := *this.connectionPtr

		for {
			n, err := c.Read(buf[:])
			if err != nil {
				this.error <- fmt.Errorf("read data error: %s", err.Error())
			}
			d := string(buf[0:n])
			if !this.ProcessCommands(d) {
				this.data <- d
			}
		}
	}
}

func (this *Client) Open() error {
	if this.state == ClientStateInit || this.state == ClientStateStopped {
		err := this.connect()
		if err != nil {
			return err
		}

		go this.readInput()
	}
	return nil
}

func (this *Client) Close() {
	if this.state == ClientStateConnected && this.connectionPtr != nil {
		c := *this.connectionPtr
		c.Close()
		this.state = ClientStateStopped
	}

}

func (this *Client) ProcessCommands(cmd string) bool {
	switch cmd {
	case CmdStop:
		this.Close()
		return true
	case CmdError:
		this.error <- fmt.Errorf("recieved error from server: %s", err.Error())
		return true
	default:
		return false
	}

	return false
}



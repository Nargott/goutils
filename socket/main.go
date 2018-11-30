package socket

import (
	"fmt"
	"encoding/json"
	"reflect"
	"strings"
	"regexp"
	"encoding/base64"
)

const (
	nameMask    = "/tmp/%s.sock"
	network     = "unixpacket"
	ReadBufSize = 512
	cmdSeparator = ":"
	cmdValidRegex = "[\\w]+"

	//protocol commands
	CmdStop  = "stop"
	CmdOk    = "ok"
	CmdError = "error"
)

func validateCmdName(cmd string) error {
	match, err := regexp.MatchString(cmdValidRegex, cmd)
	if err != nil {
		//means that bad regexp mask provided
		return err
	}
	if !match {
		return fmt.Errorf("command has bad format (not alphanumeric)")
	}

	return nil
}

func packSocketDataPacket(cmd string, data interface{}) ([]byte, error) {
	err := validateCmdName(cmd)
	if err != nil {
		return nil, err
	}

	var dataPacked []byte
	cmdPart := []byte(cmd)

	if !reflect.ValueOf(data).IsNil() { // means data present
		dataPart, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		dataPartB64 := base64.StdEncoding.EncodeToString(dataPart)

		dataPacked = append(dataPacked, cmdPart...) //add command
		dataPacked = append(dataPacked, []byte(cmdSeparator)...) //add separator
		dataPacked = append(dataPacked, dataPartB64...) //add data

		l := len(dataPacked)
		if l > ReadBufSize {
			return nil, fmt.Errorf("too long data package: %d bytes. Max len is %d bytes", l, ReadBufSize)
		}
	}

	return dataPacked, nil
}

func unpackSocketDataPacket(dataPacked []byte, receiver *interface{}) (string, error) {
	var cmd string

	dataBeginIdx := strings.Index(string(dataPacked), cmdSeparator)
	if dataBeginIdx > 0 {
		cmd = string(dataPacked[:dataBeginIdx])
		dataPart := dataPacked[(dataBeginIdx+len(cmdSeparator)):] //don't include separator

		dataPartJSON, err := base64.StdEncoding.DecodeString(string(dataPart))
		if err != nil {
			return cmd, fmt.Errorf("cannot decode data part as Base64 string: %s", err.Error())
		}

		err = json.Unmarshal(dataPartJSON, receiver)
		if err != nil {
			return cmd, fmt.Errorf("cannot decode data part as JSON: %s", err.Error())
		}
	}

	return cmd, nil
}

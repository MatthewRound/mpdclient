package mpc

import (
	"fmt"
	"os"
)


func main() {
	argsWithoutProg := os.Args[1:]
	s := mpdCommand{}
	s.port = 6600
	s.address = "192.168.0.31"
	s.parsedResult = make(map[string]string)
	command := "status"
	if argsWithoutProg[0] != "" {
		command = argsWithoutProg[0]
	}
	s.RunCommand(command)
	fmt.Printf("%s\r\n", s.rawResult)
	fmt.Printf("", s.parsedResult)
}


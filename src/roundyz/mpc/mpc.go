package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


type mpdCommand struct {
	port int32
	address string
	rawResult string
	parsedResult map[string]string
}


func (mpd *mpdCommand) ParseCommandOutput() {
	mpd.parsedResult["test"] = "test"	
	//TODO write this
	//split at end of line
	//omit first and last lines
	// break at : delimiter
}
	


func (mpd *mpdCommand) RunCommand(command string) {
	fmt.Printf("connecting to %s on port number:%d\n", mpd.address, mpd.port)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", mpd.address, mpd.port))
	if err != nil {
		fmt.Print("Cant connect\n")
	}
	cmdToSend := fmt.Sprintf("command_list_begin\n%s\ncommand_list_end\n", command)
	fmt.Fprintf(conn, cmdToSend)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		templine := scanner.Text();
		mpd.rawResult = fmt.Sprintf("%s\n%s", mpd.rawResult, templine)
		if templine == "OK" {
		  return
		}
	}
}


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


package main

import (
	"fmt"
	"bufio"
	"net"
)


type mpdCommand struct {
	port int32
	address string
	zz string
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
		mpd.zz = fmt.Sprintf("%s\n%s", mpd.zz, templine)
/* 		fmt.Print(mpd.zz) */
		if templine == "OK" {
		  return
		}
	}
}


func main() {
	s := mpdCommand{}
	s.port = 6600
	s.address = "192.168.0.31"
	s.RunCommand("currentsong")
	s.RunCommand("stats")
	fmt.Printf("%s\r\n", s.zz)
}


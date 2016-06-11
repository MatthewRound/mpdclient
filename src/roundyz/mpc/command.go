package mpc


import (
    "bufio"
    "fmt"
    "net"
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

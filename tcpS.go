package main // create tcp server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	arguments := os.Args //
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	port := ":" + arguments[1]

	// func net.Listen is listen to connection
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	//  successful call to Accept() means that the TCP server can begin to interact with TCP clients
	c, err := l.Accept() // connection
	if err != nil {
		fmt.Println(err)
		return
	}

	// we close listening when main func completed her work
	defer l.Close()

	for {
		reader := bufio.NewReader(c)
		netData, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		netDataArr := strings.Split(strings.TrimSpace(string(netData)), " ")
		command := netDataArr[0] // command entered by the client

		data := strings.Join(netDataArr[1:], " ") // data that comes after the command

		switch command {
		// stop server
		case "STOP":
			fmt.Println("Exiting TCP server!")
			return

			// returns current time in format YYYY-MM-DD HH:MM:SS WEEKDAY
		case "TIME":
			t := time.Now()
			myTime := t.Format("2006-01-02 15:04:05 Monday") + "\n"
			c.Write([]byte(myTime))

			// returns length of passed data
		case "LENGTH":
			length := strconv.Itoa(len(data)) + "\n"
			c.Write([]byte(length))

			// we entered unknown command
		default:
			c.Write([]byte("unknown command\n"))
		}

		fmt.Print("-> ", string(netData))
	}

}

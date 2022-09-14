package main // create tcp client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// we read arguments from command line
	// os.Args provides access to raw command-line arguments
	arguments := os.Args
	//fmt.Println("arguments from client", arguments)

	//checking that a value for host:port was sent
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	// we save port in CONNECT variable
	connect := arguments[1]

	// we implement of the TCP client and connect it to desired TCP server
	c, err := net.Dial("tcp", connect)
	// checking for error
	if err != nil {
		fmt.Println(err)
		return
	}

	// we create for loop to read users input from command line
	// and terminate when user send STOP command to the tcp server
	for {
		// os.Stdin allows to read data from the console
		// we create new i/o buffer reader
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CLIENT--> ")

		// we get text message, but now we ignore error
		text, _ := reader.ReadString('\n')

		// we sent text message to the TCP server over the network using Fprintf()
		fmt.Fprintf(c, text+"\n")

		// we get server response message
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("SERVER-->: " + message)

		// we terminate when user send STOP command to the tcp server
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

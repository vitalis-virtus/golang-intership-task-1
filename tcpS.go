package main // create tcp server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
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

	// we close listening when main func completed her work
	defer l.Close()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	wg := sync.WaitGroup{}

	for {
		select {
		case <-quitChan:
			fmt.Println("Stopped by command ctrl+c")
			l.Close()
			wg.Wait()
			return
		default:
		}

		c, err := l.Accept() // connection
		if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
			continue
		}

		if err != nil {
			continue
		}

		wg.Add(1)
		go func() {
			wg.Done()
			handleConnection(c)
		}()
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("connected...")
	for {

		reader := bufio.NewReader(conn)
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
			conn.Write([]byte(myTime))

			// returns length of passed data
		case "LENGTH":
			length := strconv.Itoa(len(data)) + "\n"
			conn.Write([]byte(length))

			// we entered unknown command
		default:
			conn.Write([]byte("unknown command\n"))
		}

		fmt.Print("-> ", string(netData))
	}
	conn.Close()

}

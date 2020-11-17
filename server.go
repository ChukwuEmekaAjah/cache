package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/ChukwuEmekaAjah/cacheparser"
)

type insertion struct {
	key    string
	values []string
}

var cacheMap = make(map[string]*cacheparser.KeyValue)

// var values = map[string]insertion{}

func main() {
	arguments := os.Args

	portAddress := ":1996"

	if len(arguments) > 1 {
		portAddress = ":" + arguments[1]
	}

	l, err := net.Listen("tcp", portAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go tcpHandler(c)
	}
}

func tcpHandler(c net.Conn) {

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		isValidCommand := cacheparser.Parse(strings.TrimSpace(netData))

		if !isValidCommand {
			c.Write([]byte("Invalid command sent"))
		}

		commandParts := strings.Fields(strings.TrimSpace(netData))

		commandAction, exists := cacheparser.ParserFunctions[strings.ToUpper(commandParts[0])]
		retrievalAction, ok := cacheparser.RetrievalFunctions[strings.ToUpper(commandParts[0])]

		if exists == false && ok == false {
			c.Write([]byte("Invalid command sent 2"))
		}

		if exists {
			parsedValue := commandAction(strings.ToUpper(commandParts[0]), commandParts[1:], "single")

			cacheMap[commandParts[1]] = parsedValue
		}

		// retrieval function
		if ok {
			parsedValue, err := retrievalAction(commandParts[0], commandParts[1:], cacheMap)

			if err != nil {
				c.Write([]byte("Invalid command sent 3"))
			}

			c.Write([]byte(parsedValue))
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := "\n" + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}

// ParseStdin converts client command to a map
// func parseStdin(command string) {
// 	commandParts := strings.Fields(command)

// 	if len(commandParts) < 3 {
// 		println("Commands should have action name, data identifier and data")
// 		return
// 	}

// 	commands := []string{"set", "get"}

// 	for index, command := range commands {
// 		if command == commandParts[0] {
// 			break
// 		}
// 		if index == (len(commands)-1) && command != commandParts[0] {
// 			println("Please input a valid action")
// 			return
// 		}
// 	}

// 	if commandParts[0] == "get" {
// 		fmt.Printf("value at the key is %v \n", values[commandParts[1]])
// 		return
// 	}

// 	request := new(insertion)
// 	request.key = commandParts[1]
// 	request.values = commandParts[2:]

// 	values[commandParts[1]] = *request

// 	fmt.Printf("values are %v", commandParts)
// 	println("length of commands are", len(commandParts))
// }

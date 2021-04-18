package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/rapando/gocache/entities"
	"github.com/rapando/gocache/models"
)

var memoryMap = make(map[string]*entities.DataStore)

func main() {
	port := ":7379"
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		log.Printf("Unable to listen on port %s because %v", port, err)
		return
	}
	defer listener.Close()
	log.Printf("Running on port %s", port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Printf("unable to listen : %v", err.Error())
			return
		}

		go handleConnection(connection)
	}

}

func handleConnection(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())
	var response string
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Printf("Unable to read netData because %s", err.Error())
			return
		}

		input := strings.TrimSpace(string(netData))
		if input == "" {
			c.Write([]byte("-- input is required"))
			continue
		}
		if input == "quit" {
			log.Printf("Stopping gocache...")
			break
		}

		response = router(input) + "\n"

		c.Write([]byte(response))
	}
	c.Close()
}

func router(input string) (response string) {
	command, params := processInput(input)
	switch command {
	case "save":
		location, key := models.Save(params)
		memoryMap[key] = location
		return "ok\n"

	case "all":
		log.Printf("Memory Map : %v", memoryMap)
		return "all"

	case "get":
		location := memoryMap[params[0]]
		if location == nil {
			return fmt.Sprintf("%s is missing\n", params[0])
		}
		return location.Data
	}

	return "---"
}

func processInput(input string) (command string, params []string) {
	inputArray := strings.Split(input, " ")
	log.Printf("Input Array : %v", inputArray)
	params = make([]string, len(inputArray)-1)
	command = inputArray[0]
	if len(inputArray) > 1 {
		for i := 1; i < len(inputArray); i++ {
			if inputArray[i] == " " {
				continue
			}
			params[i-1] = inputArray[i]
		}
	}
	log.Printf("Command : %s, params : %v", command, params)
	return
}

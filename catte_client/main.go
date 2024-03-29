package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Action string `json:"action"`
	Room   string `json:"room"`
	Index  int    `json:"index"`
	Data   string `json:"data"`
}

func main() {

	conn, err := net.Dial("tcp", "35.186.159.58:9999")
	if err != nil {
		fmt.Println("Server is not ready")
		os.Exit(0)
	}
	defer conn.Close()
	go run(conn)

	for {
		var com Command = Command{}
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Action: ")
		com.Action = readLine(reader)
		fmt.Print("Room: ")
		com.Room = readLine(reader)
		fmt.Print("Index: ")
		index := readLine(reader)
		if index != "" {
			fmt.Println(index)
			com.Index, _ = strconv.Atoi(index)
		}
		fmt.Print("Data: ")
		com.Data = readLine(reader)
		data, _ := json.Marshal(com)
		fmt.Println(string(data))
		fmt.Fprint(conn, string(data)+"\n")
	}
}

func readLine(reader *bufio.Reader) string {
	data, _ := reader.ReadString('\n')
	return strings.TrimSuffix(data, "\r\n")
}

func run(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()
			fmt.Println("Message from server: " + text)

			if !ok {
				fmt.Println("Reached EOF on server connection.")
				os.Exit(0)
				break
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/sarika25171/CS351-Socket-Project/model"
)

const (
	choices = "1. Create Poll\n2. List Polls\n3. Vote Poll\n4. Exit\n"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Error dialing server: ", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(choices)
		fmt.Print("Enter Choice: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading choice: ", err)
			return
		}

		switch strings.TrimSpace(choice) {
		case "1":
			fmt.Print("Enter poll question: ")

			poll := model.Poll{}
			poll.Question, err = reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading question: ", err)
				return
			}

			fmt.Print("Enter poll options (separated by commas): ")
			options, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading options: ", err)
				return
			}

			message := fmt.Sprintf("CREATE_POLL;%s:%s", strings.TrimSpace(poll.Question), strings.TrimSpace(options))
			res := SendingRequestAndReceivingResponse(conn, message)
			fmt.Println("Response from server:", res)

		case "2":
			message := "LIST_POLLS\n"
			res := SendingRequestAndReceivingResponse(conn, message)
			fmt.Println("Response from server:", res)

		case "3":
			fmt.Print("Enter poll ID to vote: ")
			pollID, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading poll ID: ", err)
				return
			}

			fmt.Print("Enter your vote: ")
			vote, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading vote: ", err)
				return
			}

			message := fmt.Sprintf("VOTE_POLL;%s:%s", strings.TrimSpace(pollID), strings.TrimSpace(vote))
			res := SendingRequestAndReceivingResponse(conn, message)
			fmt.Println("Response from server:", res)

		case "4":
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func SendingRequestAndReceivingResponse(conn net.Conn, message string) string {
	_, err := conn.Write([]byte(message + "\n"))
	if err != nil {
		log.Println("Error sending message: ", err)
		return ""
	}

	response, err := bufio.NewReader(conn).ReadString(';')
	if err != nil {
		log.Println("Error reading response: ", err)
		return ""
	}
	result := strings.SplitN(response, ";", 2)

	return result[0]
}

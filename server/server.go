package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/sarika25171/CS351-Socket-Project/model"
)

const (
	transportLayerProtocol = "tcp"
	port                   = "8080"
)

var (
	polls   = make(map[string]model.Poll)
	idMutex = sync.Mutex{}
)

func main() {
	InitServer()
}

func InitServer() {
	server, err := net.Listen(transportLayerProtocol, fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
	defer server.Close()
	log.Printf("Server is running on port %s\n", port)

	for {
		connection, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(connection.RemoteAddr().String(), "connected")
		go HandleConnection(connection)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading message: ", err)
			return
		}

		message = strings.TrimSpace(message)
		log.Printf("Message from client '%s' : %s\n", conn.RemoteAddr().String(), message)

		response := processMessage(message)
		conn.Write([]byte(response + "\n"))
	}
}

func processMessage(message string) string {
	parts := strings.SplitN(message, ";", 2)
	if len(parts) < 2 && parts[0] != "LIST_POLLS" {
		return "[400] Invalid message format\n;"
	}

	command := parts[0]
	var data string
	if command != "LIST_POLLS" {
		data = parts[1]
	}

	switch command {
	case "CREATE_POLL":
		return handleCreatePoll(data)
	case "LIST_POLLS":
		return handleListPolls()
	case "VOTE_POLL":
		return handleVotePoll(data)
	default:
		return "[401] Unknown command\n;"
	}
}

func handleCreatePoll(data string) string {
	lines := strings.SplitN(data, ":", 2)
	if len(lines) < 2 {
		return "[402] Invalid format for CREATE_POLL\n;"
	}

	question := lines[0]
	options := strings.Split(lines[1], ",")

	idMutex.Lock()
	defer idMutex.Unlock()

	pollID := fmt.Sprintf("%d", len(polls)+1)
	votes := make(map[string]int)
	for _, option := range options {
		votes[option] = 0
	}
	polls[pollID] = model.Poll{
		ID:       pollID,
		Question: question,
		Options:  options,
		Votes:    votes,
	}

	return fmt.Sprintf("[201] Poll created successfully!\nID: %s\nQuestion: %s\nOptions: %s\n;", pollID, question, strings.Join(options, ","))
}

func handleListPolls() string {
	if len(polls) == 0 {
		return "[300] No polls available;"
	}

	var result string
	for _, poll := range polls {
		result += fmt.Sprintf("\n\nID: %s\nQuestion: %s\nOptions: %s\nVotes: %v\n\n", poll.ID, poll.Question, strings.Join(poll.Options, ","), poll.Votes)
	}

	return "\n[200] Poll created successfully!" + result + ";"
}

func handleVotePoll(data string) string {
	lines := strings.SplitN(data, ":", 2)
	if len(lines) < 2 {
		return "[403] Invalid format for VOTE_POLL\n;"
	}

	pollID := lines[0]
	vote := lines[1]

	poll, exists := polls[pollID]
	if !exists {
		return "[404] Poll ID not found\n;"
	}

	if !contains(poll.Options, vote) {
		return "[405] Invalid vote option\n;"
	}

	poll.Votes[vote]++
	polls[pollID] = poll

	return fmt.Sprintf("[203] Vote recorded for poll ID %s: %s\n;", pollID, vote)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

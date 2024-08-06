# Voting Poll with Socket Programming

## Introduction
This project implements a client-server application for managing polls. The client (`client.go`) allows users to create polls, list existing polls, and vote on them interactively through a TCP connection with the server (`server.go`).

## Key Features
- **Create Poll**: Users can create a new poll by entering a question and a list of options.
- **List Polls**: Users can view a list of existing polls stored on the server.
- **Vote Poll**: Users can vote on a specific poll by providing the poll ID and their vote choice.
- **Client-Server Communication**: Communication between the client and server is handled over TCP/IP sockets.

## Project Structure
- **client.go**: Implements the client-side functionality.
- **server.go**: Implements the server-side functionality.
- **model**: Contains data models used for representing polls.

## Usage
To use this application:
1. Start the server using `go run server.go`.
2. Start the client using `go run client.go`.
3. Follow the prompts to create polls, list existing polls, and vote.

## Contributors
- Created by Sarika Opastpitchayadamrong.

## Protocol Defined As
Communication between the client and server follows a simple protocol:
- **CREATE_POLL**: Creates a new poll with a specified question and options.
- **LIST_POLLS**: Requests a list of all available polls.
- **VOTE_POLL**: Registers a vote for a specified poll ID.

## Status Codes
- **200**: Poll listed successfully 
- **201**: Poll created successfully
- **300**: No polls available
- **400**: Invalid message format
- **401**: Unknown command
- **402**: Invalid format for CREATE_POLL
- **403**: Invalid format for VOTE_POLL
- **404**: Poll ID not found
- **405**: Invalid vote option

## Procedure of Working
1. **Client Operation**: Prompts the user for choices (create, list, vote).
2. **Server Operation**: Handles client requests, processes data, and sends responses.
3. **Communication**: Uses TCP sockets for reliable communication.

## Instructions
- Ensure Go (Golang) is installed on your machine.
- Run `go run server.go` to start the server.
- Run `go run client.go` to start the client.
- Follow on-screen prompts to interact with the application.

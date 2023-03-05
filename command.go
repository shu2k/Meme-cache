package main

import "time"

type Command string

const (
	CMDSet Command = "SET"
	CMDGet Command = "GET"
)

type Message struct {
	Cmd   Command
	Key   []byte
	Valye []byte
	TTL   time.Duration
}

func parseCommand(buffer []byte) (*Message, error) {

}

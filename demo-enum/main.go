package main

import "fmt"

type MessageType uint8

const (
	GPB MessageType = iota
	COAP
	JSON
)

func main() {
	types := MessageType(JSON)
	switch types {
	case GPB:
		fmt.Println("GPB")
	case COAP:
		fmt.Println("COAP")
	case JSON:
		fmt.Println("JSON")
	}
}

package main

import "rabbitmq/stream"

func main() {
	stream.Send("This is the message")
	stream.Receive()
}

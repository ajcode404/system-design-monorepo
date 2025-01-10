package main

import (
	"bufio"
	"os"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func main() {
	env, err := stream.NewEnvironment(stream.NewEnvironmentOptions().SetHost("localhost").SetPort(5552))
	if err != nil {
		panic(err)
	}
	streamName := "hello-go-stream"
	env.DeclareStream(streamName, &stream.StreamOptions{
		MaxLengthBytes: stream.ByteCapacity{}.GB(2),
	})

	producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
	if err != nil {
		panic(err)
	}
	err = producer.Send(amqp.NewMessage([]byte("Hi AJ")))
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	producer.Close()
}

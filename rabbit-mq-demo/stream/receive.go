package stream

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func Receive() {
	env, err := stream.NewEnvironment(stream.NewEnvironmentOptions())
	if err != nil {
		panic(err)
	}
	streamName := "hello-go-stream"
	env.DeclareStream(streamName, &stream.StreamOptions{
		MaxLengthBytes: stream.ByteCapacity{}.GB(2),
	})
	messageHandler := func(cc stream.ConsumerContext, message *amqp.Message) {
		fmt.Printf("Stream %s and Received message: %s\n", cc.Consumer.GetStreamName(), message.Data)
	}
	consume, err := env.NewConsumer(streamName, messageHandler, stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to exit")
	_, _ = reader.ReadString('\n')
	consume.Close()
}

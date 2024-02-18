package messagequeue

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func MQCreateConsumer() {
	con, err := amqp091.Dial("amqp://user:mrc201@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer con.Close()

	channel, err := con.Channel()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	consumer, err := channel.ConsumeWithContext(ctx, "create", "create-consumer", true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for msg := range consumer {
		fmt.Println("ROUTING_KEY: ", msg.RoutingKey)
		fmt.Println("Body: ", string(msg.Body))
	}
}

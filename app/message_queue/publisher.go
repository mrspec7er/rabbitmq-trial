package messagequeue

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func MQCreatePublisher() {
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
	message := amqp091.Publishing{
		ContentType: "application/json",
		Body:        []byte("Hello From Publisher"),
	}

	err = channel.PublishWithContext(ctx, "report", "create", false, false, message)

	if err != nil {
		fmt.Println(err)
	}
}

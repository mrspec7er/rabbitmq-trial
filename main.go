package main

import messagequeue "github.com/mrspec7er/rabbitmq-trials/app/message_queue"

func main() {
	messagequeue.MQCreateConsumer()
}

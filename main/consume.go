package main

import (
	"fmt"
	randomstrings "github.com/hekonsek/random-strings"
	"github.com/spf13/cobra"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func init() {
	RootCommand.AddCommand(ConsumeCommand)
}

var ConsumeCommand = &cobra.Command{
	Use:   "consume",
	Short: "consume",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
			"group.id":          randomstrings.ForHumanWithDashAndHash(),
			"auto.offset.reset": "earliest",
		})

		if err != nil {
			panic(err)
		}

		c.SubscribeTopics([]string{"myTopic"}, nil)

		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}

		c.Close()
	},
}

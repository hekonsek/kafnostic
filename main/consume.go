package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	randomstrings "github.com/hekonsek/random-strings"
	"github.com/spf13/cobra"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var consumeCluster string

func init() {
	consumeCommand.Flags().StringVarP(&consumeCluster, "cluster", "", "localhost", "")
	RootCommand.AddCommand(consumeCommand)
}

var consumeCommand = &cobra.Command{
	Use:   "consume",
	Short: "consume",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": consumeCluster,
			"group.id":          randomstrings.ForHumanWithDashAndHash(),
			"auto.offset.reset": "earliest",
		})
		defer func() {
			osexit.ExitOnError(c.Close())
		}()

		if err != nil {
			panic(err)
		}

		osexit.ExitOnError(c.SubscribeTopics([]string{"myTopic"}, nil))

		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	},
}

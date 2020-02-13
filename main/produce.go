package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	randomstrings "github.com/hekonsek/random-strings"
	"github.com/spf13/cobra"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)

var produceCluster string

func init() {
	ProduceCommand.Flags().StringVarP(&produceCluster, "cluster", "", "localhost", "")
	RootCommand.AddCommand(ProduceCommand)
}

var ProduceCommand = &cobra.Command{
	Use:   "produce",
	Short: "produce",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": produceCluster})
		if err != nil {
			osexit.ExitOnError(err)
		}

		defer p.Close()

		// Delivery report handler for produced messages
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()

		topic := "myTopic"
		for {
			payload := fmt.Sprintf(`{"%s": "%s"}`, randomstrings.ForHumanWithDashAndHash(), randomstrings.ForHumanWithDashAndHash())
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(payload),
			}, nil)
			time.Sleep(time.Millisecond * 200)
		}
	},
}

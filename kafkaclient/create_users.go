package main

import (
	"fmt"

	"awesome-go/kafkaclient/cmd"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/twmb/franz-go/pkg/kmsg"
)

func main() {
	kgoOpts, err := NewKgoConfig(&Kafka{
		Brokers:  []string{"127.0.0.1:19095", "127.0.0.1:19093", "127.0.0.1:19094"},
		ClientID: "client-2",
		SASL: KafkaSASL{
			Mechanism: SASLMechanismScramSHA256,
			Username:  "admin",
			Password:  "kafka123456",
		},
	})
	if err != nil {
		fmt.Errorf("failed to create a valid kafka client config: %w", err)
	}

	kafkaClient, err := kgo.NewClient(kgoOpts...)
	if err != nil {
		fmt.Errorf("failed to create kafka client: %w", err)
	}

	// 3. Submit reassign partitions request
	kmsgReq := make([]kmsg.AlterPartitionAssignmentsRequestTopic, 0)
	partitions := make([]kmsg.AlterPartitionAssignmentsRequestTopicPartition, 2)

	partitionReq := kmsg.NewAlterPartitionAssignmentsRequestTopicPartition()
	partitionReq.Partition = 0
	partitionReq.Replicas = []int32{2, 1}
	partitions = append(partitions, partitionReq)

	partitionReq2 := kmsg.NewAlterPartitionAssignmentsRequestTopicPartition()
	partitionReq2.Partition = 1
	partitionReq2.Replicas = []int32{2, 1}
	partitions = append(partitions, partitionReq2)

	topicReq := kmsg.NewAlterPartitionAssignmentsRequestTopic()
	topicReq.Topic = "test-5"
	topicReq.Partitions = partitions
	kmsgReq = append(kmsgReq, topicReq)

	cmd.Rebalanced(kafkaClient, kmsgReq)

}

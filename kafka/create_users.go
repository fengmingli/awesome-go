package main

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/twmb/franz-go/pkg/kmsg"
)

func main() {
	CreateUsers(context.TODO())
}

// CreateUsers creates a Kafka topic.
func CreateUsers(ctx context.Context) (*[]kmsg.AlterUserSCRAMCredentialsResponseResult, error) {
	kgoOpts, err := NewKgoConfig(&Kafka{
		Brokers:  []string{"127.0.0.1:19092", "127.0.0.1:19093", "127.0.0.1:19094"},
		ClientID: "ddddddddd",
		SASL: KafkaSASL{
			Mechanism: SASLMechanismScramSHA256,
			Username:  "admin",
			Password:  "admin123456",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create a valid kafka client config: %w", err)
	}

	kafkaClient, err := kgo.NewClient(kgoOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka client: %w", err)
	}

	req := kmsg.NewAlterUserSCRAMCredentialsRequest()
	req.Upsertions = []kmsg.AlterUserSCRAMCredentialsRequestUpsertion{
		{
			Name:      "lfm-eeeee",
			Mechanism: 1,
			//Iterations:     4096,
			Salt:           []byte("xxxxxxx"),
			SaltedPassword: []byte("xxxxxxx"),
		},
	}
	res, err := req.RequestWith(ctx, kafkaClient)
	if err != nil {
		return nil, fmt.Errorf("request has failed: %w", err)
	}
	if len(res.Results) != 1 {
		return nil, fmt.Errorf("unexpected number of topic responses, expected exactly one but got '%v'", len(res.Results))
	}

	return &res.Results, nil
}

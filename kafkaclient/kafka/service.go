// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package kafka

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

// Service acts as interface to interact with the Kafka Cluster
type Service struct {
	KafkaClient *kgo.Client
}

// NewService creates a new Kafka service and immediately checks connectivity to all components. If any of these external
// dependencies fail an error wil be returned.
func NewService(client *kgo.Client) (*Service, error) {
	return &Service{
		KafkaClient: client,
	}, nil
}

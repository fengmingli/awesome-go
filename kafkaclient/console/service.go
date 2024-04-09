// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package console

import (
	"awesome-go/kafkaclient/kafka"
)

// Service offers all methods to serve the responses for the REST API. This usually only involves fetching
// several responses from Kafka concurrently and constructing them so, that they are
type Service struct {
	kafkaSvc *kafka.Service
}

// NewService for the Console package
func NewService(
	kafkaSvc *kafka.Service,
) (*Service, error) {
	return &Service{
		kafkaSvc: kafkaSvc,
	}, nil
}

// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package main

import (
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kversion"
	"github.com/twmb/franz-go/pkg/sasl"
	"github.com/twmb/franz-go/pkg/sasl/scram"
)

// NewKgoConfig creates a new Config for the Kafka Client as exposed by the franz-go library.
// If TLS certificates can't be read an error will be returned.
//
//nolint:gocognit,cyclop // This function is lengthy, but it's only plumbing configurations. Seems okay to me.
func NewKgoConfig(cfg *Kafka) ([]kgo.Opt, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.MaxVersions(kversion.V2_7_0()),
		kgo.ClientID(cfg.ClientID),
		kgo.FetchMaxBytes(5 * 1000 * 1000), // 5MB
		kgo.MaxConcurrentFetches(12),
		// We keep control records because we need to consume them in order to know whether the last message in a
		// a partition is worth waiting for or not (because it's a control record which we would never receive otherwise)
		kgo.KeepControlRecords(),
	}

	// SASL SCRAM
	if cfg.SASL.Mechanism == SASLMechanismScramSHA256 || cfg.SASL.Mechanism == SASLMechanismScramSHA512 {
		var mechanism sasl.Mechanism
		scramAuth := scram.Auth{
			User: cfg.SASL.Username,
			Pass: cfg.SASL.Password,
		}
		if cfg.SASL.Mechanism == SASLMechanismScramSHA256 {
			mechanism = scramAuth.AsSha256Mechanism()
		}
		if cfg.SASL.Mechanism == SASLMechanismScramSHA512 {
			mechanism = scramAuth.AsSha512Mechanism()
		}
		opts = append(opts, kgo.SASL(mechanism))
	}

	return opts, nil
}

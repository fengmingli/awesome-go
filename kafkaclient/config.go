// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

// Package config defines all Configuration options for Console. It also contains
// the parsing and validation logic.
package main

// Config holds all (subdependency)Configs needed to run the API
type Config struct {
	ConfigFilepath   string
	MetricsNamespace string `yaml:"metricsNamespace"`
	ServeFrontend    bool   `yaml:"serveFrontend"` // useful for local development where we want the frontend from 'npm run start'

	Kafka Kafka `yaml:"kafka"`
}

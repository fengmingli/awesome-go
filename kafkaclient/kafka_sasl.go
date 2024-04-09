// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package main

const (
	// SASLMechanismPlain represents the configuration value for SASL Plain authentication.
	SASLMechanismPlain = "PLAIN"
	// SASLMechanismScramSHA256 represents the configuration value for SASL scram-sha-256 authentication.
	SASLMechanismScramSHA256 = "SCRAM-SHA-256"
	// SASLMechanismScramSHA512 represents the configuration value for SASL scram-sha-512 authentication.
	SASLMechanismScramSHA512 = "SCRAM-SHA-512"
	// SASLMechanismGSSAPI represents the configuration value for SASL Kerberos authentication.
	SASLMechanismGSSAPI = "GSSAPI"
	// SASLMechanismOAuthBearer represents the configuration value for SASL OAuth authentication.
	SASLMechanismOAuthBearer = "OAUTHBEARER"
	// SASLMechanismAWSManagedStreamingIAM represents the configuration value for authenticating via AWS MSK IAM.
	SASLMechanismAWSManagedStreamingIAM = "AWS_MSK_IAM"
)

// KafkaSASL for Kafka client
type KafkaSASL struct {
	Enabled   bool   `yaml:"enabled"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Mechanism string `yaml:"mechanism"`
}

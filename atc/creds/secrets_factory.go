package creds

import (
	"time"
)

//counterfeiter:generate . SecretsFactory
type SecretsFactory interface {
	// NewSecrets returns an instance of a secret manager, capable of retrieving individual secrets
	NewSecrets() Secrets
}

//counterfeiter:generate . Secrets
type Secrets interface {
	// Every credential manager needs to be able to return (secret, secret_expiration_time, exists, error) based on the secret path
	Get(string) (any, *time.Time, bool, error)

	// NewSecretLookupPaths returns an instance of lookup policy, which can transform pipeline ((var)) into one or more secret paths, based on team name and pipeline name
	NewSecretLookupPaths(string, string, bool) []SecretLookupPath
}

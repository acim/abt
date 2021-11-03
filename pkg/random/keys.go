package random

import (
	"fmt"
)

const (
	accessKeyChars       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secretAccessKeyChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
)

// GenerateKeys generates AWS compatible access and secret access keys.
func GenerateKeys() (accessKey string, secretAccessKey string, err error) {
	accessKey, err = Random(accessKeyChars, 20) //nolint:gomnd
	if err != nil {
		return "", "", fmt.Errorf("random access key: %w", err)
	}

	sak, err := Random(secretAccessKeyChars, 40) //nolint:gomnd
	if err != nil {
		return "", "", fmt.Errorf("random secret access key: %w", err)
	}

	return accessKey, sak, nil
}

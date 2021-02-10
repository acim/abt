package aws

import (
	"fmt"

	"github.com/acim/abt/pkg/strings"
)

const (
	accessKeyChars       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secretAccessKeyChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
)

// GenerateKeys generates AWS compatible access and secret access keys.
func GenerateKeys() (accessKey string, secretAccessKey string, err error) {
	ak, err := strings.Random(accessKeyChars, 20)
	if err != nil {
		return "", "", fmt.Errorf("random access key: %w", err)
	}

	sak, err := strings.Random(secretAccessKeyChars, 40)
	if err != nil {
		return "", "", fmt.Errorf("random secret access key: %w", err)
	}

	return ak, sak, nil
}

package aws

import (
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
		return "", "", err
	}
	sak, err := strings.Random(secretAccessKeyChars, 40)
	if err != nil {
		return "", "", err
	}
	return ak, sak, nil
}

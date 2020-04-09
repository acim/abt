package strings

import (
	"crypto/rand"
	"errors"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Random generates random string of provided character set and length.
func Random(availableChars string, length int) (string, error) {
	availableCharLength := len(availableChars)
	if availableCharLength == 0 || availableCharLength > 256 {
		panic("availableChars length must be greater than 0 and less than or equal to 256")
	}
	var bitLength byte
	var bitMask byte
	for bits := availableCharLength - 1; bits != 0; {
		bits = bits >> 1
		bitLength++
	}
	bitMask = 1<<bitLength - 1

	bufferSize := length + length/3
	var err error

	result := make([]byte, length)
	for i, j, rb := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			rb, err = randomBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(rb[j%length] & bitMask); idx < availableCharLength {
			result[i] = availableChars[idx]
			i++
		}
	}

	return string(result), nil
}

func randomBytes(length int) ([]byte, error) {
	var rb = make([]byte, length)
	_, err := rand.Read(rb)
	if err != nil {
		return nil, errors.New("unable to generate random bytes")
	}
	return rb, nil
}
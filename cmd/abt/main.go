package main

import (
	"fmt"
	"os"

	abt "github.com/acim/abt/pkg/random"
)

func main() {
	accessKey, sak, err := abt.GenerateKeys()
	if err != nil {
		fmt.Println(err) //nolint:forbidigo
		os.Exit(1)
	}

	fmt.Println("Access key:", accessKey)  //nolint:forbidigo
	fmt.Println("Secret access key:", sak) //nolint:forbidigo
}

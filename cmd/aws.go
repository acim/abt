package cmd

import (
	"fmt"

	"github.com/acim/tools/pkg/aws"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws [genkeys]",
	Short: "AWS related command subset",
}

var awsGenKeysCmd = &cobra.Command{
	Use:   "genkeys",
	Short: "Generate AWS compatible access and access secret keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		ak, sak, err := aws.GenerateKeys()
		if err != nil {
			return err
		}
		fmt.Println("Access key:", ak)
		fmt.Println("Secret access key:", sak)
		return nil
	},
}

func init() {
	awsCmd.AddCommand(awsGenKeysCmd)
	rootCmd.AddCommand(awsCmd)
}

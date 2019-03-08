package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `test PersistentFlags`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(userLicense)
		fmt.Println(auth)
	},
}
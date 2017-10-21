package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// version command to tell ecom version
var version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number for ecom",
	Long:  `Build version of ecom binary`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ecom version: 0.0.1")
	},
}

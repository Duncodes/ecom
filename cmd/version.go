package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "print the version number for ecom",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ecom version: 0.0.1")
	},
}
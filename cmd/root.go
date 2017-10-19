package cmd

import "github.com/spf13/cobra"

// RootCmd ....
var RootCmd = &cobra.Command{
	Use:   "ecom",
	Short: "Ecom is a ecommerce api ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

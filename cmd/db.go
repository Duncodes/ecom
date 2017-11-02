package cmd

import "github.com/spf13/cobra"

func init() {
	dbcmd.AddCommand(drop)
}

var dbcmd = &cobra.Command{
	Use:   "db",
	Short: "Manages ecom databases",
	Long:  `The manages ecom database like a pro`,
}

var drop = &cobra.Command{
	Use:   "drop",
	Short: "drops all tables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO run remove database command
		//database.DropTables()
	},
}

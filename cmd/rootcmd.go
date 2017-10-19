package cmd

import (
	"log"

	"github.com/Duncodes/ecom/config"
	"github.com/Duncodes/ecom/database"
	"github.com/Duncodes/ecom/server"
	"github.com/spf13/cobra"
)

var port string

var configfilepath string

func loadconfig() {
	err := config.LoadConfig(configfilepath)
	if err != nil {
		log.Println(err)
		log.Println("Falling back to default configuration")
	}

}
func init() {
	serveCmd.PersistentFlags().StringVarP(&config.Config.Port, "port", "p", "9200", "Port for the server")

	RootCmd.PersistentFlags().StringVarP(&configfilepath, "config", "f", "config.json", "Ecom Config json file")
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(version)

}

// RootCmd ....
var RootCmd = &cobra.Command{
	Use:   "ecom",
	Short: "Ecom is a ecommerce api ",
	Long:  `Ecom a less usefull ecommerce api server for anyone who loves to code`,
	Run: func(cmd *cobra.Command, args []string) {
		loadconfig()
		database.InitDB()

		/* get the port to use . if flag is provide use flag else use config
		file
		*/

		server.StartServer(config.Config.Port)
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts ecom server",
	Long: `serve starts the ecom http serve ,
			by default is uses port :9200 port can be provides via a flag,`,
	Run: func(cmd *cobra.Command, args []string) {
		loadconfig()
		database.InitDB()

		server.StartServer(config.Config.Port)

	},
}

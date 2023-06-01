package server

import (
	"github.com/golang-training-examples/example2/cmd/root"
	"github.com/golang-training-examples/example2/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run HTTP server",
	Aliases: []string{"s"},
	Run: func(c *cobra.Command, args []string) {
		server.Server(FlagPort)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)

	viper.BindEnv("SERVER_PORT")
	port := viper.GetInt("SERVER.PORT")
	portEnv := viper.GetInt("SERVER_PORT")
	if portEnv != 0 {
		port = portEnv
	}

	Cmd.Flags().IntVarP(
		&FlagPort,
		"port",
		"p",
		port,
		"Listen on port",
	)
	if port == 0 {
		Cmd.MarkFlagRequired("port")
	}
}

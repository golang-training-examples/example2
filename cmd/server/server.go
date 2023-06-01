package server

import (
	"github.com/golang-training-examples/example2/cmd/root"
	"github.com/golang-training-examples/example2/pkg/server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run HTTP server",
	Aliases: []string{"s"},
	Run: func(c *cobra.Command, args []string) {
		server.Server(8000)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}

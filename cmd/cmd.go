package cmd

import (
	"github.com/golang-training-examples/example2/cmd/root"
	_ "github.com/golang-training-examples/example2/cmd/server"
	_ "github.com/golang-training-examples/example2/cmd/status"
	_ "github.com/golang-training-examples/example2/cmd/version"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}

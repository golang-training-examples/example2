package cmd

import (
	"github.com/golang-training-examples/example2/cmd/root"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}

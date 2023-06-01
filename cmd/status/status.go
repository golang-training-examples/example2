package status

import (
	"github.com/golang-training-examples/example2/cmd/root"
	"github.com/golang-training-examples/example2/pkg/status"
	"github.com/spf13/cobra"
)

var FlagJSON bool

var Cmd = &cobra.Command{
	Use:     "status",
	Short:   "Get server status",
	Aliases: []string{"st"},
	Run: func(c *cobra.Command, args []string) {
		if FlagJSON {
			status.PrintStatusJsonOrDie()
		} else {
			status.PrintStatusOrDie()
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagJSON,
		"json",
		false,
		"Show output in JSON",
	)
}
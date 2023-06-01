package status

import (
	"github.com/golang-training-examples/example2/cmd/root"
	"github.com/golang-training-examples/example2/pkg/status"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagJSON bool
var FlagOrigin string

var Cmd = &cobra.Command{
	Use:     "status",
	Short:   "Get server status",
	Aliases: []string{"st"},
	Run: func(c *cobra.Command, args []string) {
		if FlagJSON {
			status.PrintStatusJsonOrDie(FlagOrigin)
		} else {
			status.PrintStatusOrDie(FlagOrigin)
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

	viper.BindEnv("CLIENT_ORIGIN")
	origin := viper.GetString("CLIENT.ORIGIN")
	originEnv := viper.GetString("CLIENT_ORIGIN")
	if originEnv != "" {
		origin = originEnv
	}

	Cmd.Flags().StringVar(
		&FlagOrigin,
		"origin",
		origin,
		"Connect to server on origin",
	)
	if origin == "" {
		Cmd.MarkFlagRequired("origin")
	}
}

package root

import (
	"fmt"

	"github.com/ondrejsika/go-hello"
	"github.com/spf13/cobra"
	gitlab_hello "gitlab.sikalabs.com/go/hello-from-gitlab"
)

var FlagName string

var Cmd = &cobra.Command{
	Use: "example2",
	Run: func(c *cobra.Command, args []string) {
		hello.Hello()
		gitlab_hello.HelloFromGitlab()
		fmt.Println("Hello", FlagName)
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}

package main

import (
	"fmt"

	"github.com/ondrejsika/go-hello"
	gitlab_hello "gitlab.sikalabs.com/go/hello-from-gitlab"
)

func main() {
	hello.Hello()
	gitlab_hello.HelloFromGitlab()
	fmt.Println("Hello World")
}

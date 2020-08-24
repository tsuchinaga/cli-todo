package main

import "gitlab.com/tsuchinaga/cli-todo/infra/cli"

func main() {
	if err := cli.Run(); err != nil {
		panic(err)
	}
}

package main

import "github.com/GOGOsu/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}

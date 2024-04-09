package main

import (
	"context"
	"flag"
	"os"

	"github.com/Akshat120/coders/CLI/base32"
	"github.com/Akshat120/coders/CLI/base64"
	"github.com/google/subcommands"
)

func main() {
	os.Exit(run())
}

func run() int {

	groupLabel := "base encoder"
	subcommands.Register(&base64.Command{}, groupLabel)
	subcommands.Register(&base32.Command{}, groupLabel)

	flag.Parse()
	status := subcommands.Execute(context.Background())

	return int(status)
}

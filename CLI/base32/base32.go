package base32

import (
	"context"
	"encoding/base32"
	"flag"
	"fmt"

	"github.com/dcowgill/envflag"
	"github.com/google/subcommands"
)

type Command struct {
	mode  string
	input string
}

func (*Command) Name() string { return "base32" }

func (*Command) Synopsis() string { return "Performs base32 encode/decode" }

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.mode, "mode", "", "mode can be encode or decode")
	f.StringVar(&c.input, "input", "", "input is a string")
}

func (*Command) Usage() string {
	return `USAGE base32 -mode <mode> -input <input>

Performs base32 encode/decode

Options:
`
}

func (c *Command) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := envflag.NewVarSet(f).Parse(); err != nil {
		fmt.Printf("error parsing flags: %v\n", err)
		return subcommands.ExitFailure
	}

	output := ""

	if c.mode == "encode" {
		output = base32.StdEncoding.EncodeToString([]byte(c.input))
	} else if c.mode == "decode" {
		outputBytes, err := base32.StdEncoding.DecodeString(c.input)
		if err != nil {
			fmt.Println("Error:", err)
			return subcommands.ExitFailure
		}
		output = string(outputBytes)
	} else {
		return subcommands.ExitFailure
	}

	fmt.Println("Output:", output)

	return subcommands.ExitSuccess
}

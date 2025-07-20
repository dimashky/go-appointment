package main

import (
	"fmt"
	"os"

	"github.com/dimashky/go-appointment/app/commands"
)

const Version = "0.0.1"

func main() {
	if err := commands.Execute(Version); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/Zac-Garby/plp/args"
	"github.com/Zac-Garby/plp/operators"
)

func main() {
	arg := os.Args
	if len(arg) <= 1 {
		fmt.Println(`plp - Pluto Package Manager

  +<package>	installs a package
  -<package>	removes a package
  ^<package>	updates a package`)

		return
	}

	ops, err := args.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, op := range ops {
		var err error

		switch op.Type {
		case args.INSTALL:
			err = operators.Install(op.Package)
		case args.REMOVE:
			err = operators.Remove(op.Package)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

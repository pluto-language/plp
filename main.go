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
		fmt.Println(`  < plp - Pluto Package Manager >

  +<package>	installs a package
  -<package>	removes a package
  ^<package>	updates a package
`)

		return
	}

	ops, err := args.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		installs = 0
		removes  = 0
		updates  = 0
	)

	for _, op := range ops {
		var err error

		switch op.Type {
		case args.INSTALL:
			err = operators.Install(op.Package)
			installs++
		case args.REMOVE:
			err = operators.Remove(op.Package)
			removes++
		case args.UPDATE:
			err = operators.Update(op.Package)
			updates++
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Printf("\n%d package(s) installed\n%d package(s) removed\n%d package(s) updated\n", installs, removes, updates)
}

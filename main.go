package main

import (
	"fmt"
	"os"

	"github.com/Zac-Garby/plp/args"
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

	fmt.Println(ops)
}

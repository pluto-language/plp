package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

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
  *<package>    generates a package
`)

		return
	}

	if arg[1] == "list" {
		err := operators.List()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return
	}

	ops, err := args.Parse(arg[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		installs  = 0
		removes   = 0
		updates   = 0
		generates = 0
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

	fmt.Println()

	var (
		green = color.New(color.FgGreen, color.Bold)
		cyan  = color.New(color.FgCyan, color.Bold)
		red   = color.New(color.FgRed, color.Bold)
	)

	printStat(green, "INSTALLED", installs)
	printStat(red, "REMOVED", removes)
	printStat(cyan, "UPDATED", updates)
	printStat(cyan, "GENERATED", generates)
}

func printStat(colour *color.Color, prefix string, count int) {
	if count > 0 {
		colour.Printf("%-10s ", prefix)

		if count > 1 {
			fmt.Printf("%d packages", count)
		} else {
			fmt.Printf("1 package")
		}

		fmt.Println()
	}
}

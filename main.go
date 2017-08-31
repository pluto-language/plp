package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/Zac-Garby/plp/args"
	"github.com/Zac-Garby/plp/operators"
)

func main() {
	arg := os.Args
	if len(arg) <= 1 {
		fmt.Println(`  < plp - Pluto Package Manager >

  plp +<package>	installs a package
  plp -<package>	removes a package
  plp ^<package>	updates a package

  plp list          lists all available packages
  plp gen           creates a new package
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

	if arg[1] == "gen" {
		err := operators.Generate()

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
		installs []string
		removes  []string
		updates  []string
	)

	for _, op := range ops {
		var err error

		switch op.Type {
		case args.INSTALL:
			err = operators.Install(op.Package)
			installs = append(installs, op.Package)
		case args.REMOVE:
			err = operators.Remove(op.Package)
			removes = append(removes, op.Package)
		case args.UPDATE:
			err = operators.Update(op.Package)
			updates = append(updates, op.Package)
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
}

func printStat(colour *color.Color, prefix string, pkgs []string) {
	count := len(pkgs)

	if count > 0 {
		colour.Printf("%10s ", prefix)

		if count > 1 {
			fmt.Printf("%d packages", count)
		} else {
			fmt.Printf("1 package")
		}

		fmt.Printf(" (%s)\n", strings.Join(pkgs, ", "))
	}
}

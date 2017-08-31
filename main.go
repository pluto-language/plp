package main

import "github.com/Zac-Garby/plp/args"
import "os"
import "fmt"

func main() {
	ops, err := args.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ops)
}

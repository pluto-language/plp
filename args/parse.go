package args

import (
	"errors"
	"fmt"
)

type opType string

// Operation types
const (
	// INSTALL - Installs a package
	INSTALL opType = "install"

	// REMOVE - Removes an installed package
	REMOVE opType = "remove"

	// UPDATE - Updates a package to the latest version
	UPDATE opType = "update"
)

// Operation represents an operation, such as install
// takes one argument: the package name.
type Operation struct {
	Type    opType
	Package string
}

func (o Operation) String() string {
	return fmt.Sprintf("%s %s", o.Type, o.Package)
}

// Parse parses a slice of command line arguments
// into a slice of operations
func Parse(args []string) ([]Operation, error) {
	var ops []Operation

	types := map[byte]opType{
		'+': INSTALL,
		'-': REMOVE,
		'^': UPDATE,
	}

	for _, arg := range args {
		first := arg[0]

		if len(arg) <= 1 {
			msg := fmt.Sprintf("arg error: a package name is expected after '%s'", string(first))
			return ops, errors.New(msg)
		}

		if ty, ok := types[first]; ok {
			op := Operation{
				Type:    ty,
				Package: arg[1:],
			}

			ops = append(ops, op)
		} else {
			msg := fmt.Sprintf("arg error: unknown operator type '%s'", string(first))
			return ops, errors.New(msg)
		}
	}

	return ops, nil
}

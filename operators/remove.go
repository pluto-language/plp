package operators

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Remove removes an installed package
func Remove(pkg string) error {
	root, err := getRoot()
	if err != nil {
		return err
	}

	path := filepath.Join(root, "libraries", pkg)

	if _, err := os.Stat(path); err != nil {
		msg := fmt.Sprintf("package '%s' not found at '%s'", pkg, path)
		return errors.New(msg)
	}

	err = os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}

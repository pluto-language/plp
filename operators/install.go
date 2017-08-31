package operators

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
)

// Install downloads and installs a
// package named 'pkg'
func Install(pkg string) error {
	downloadDirectory()

	packages, err := getPackages()
	if err != nil {
		return err
	}

	repo, ok := packages[pkg]
	if !ok {
		msg := fmt.Sprintf("package '%s' not found in $PLUTO/libraries/packages.json", pkg)
		return errors.New(msg)
	}

	root, err := getRoot()
	if err != nil {
		return err
	}

	_, err = git.PlainClone(filepath.Join(root, "libraries", pkg), false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}

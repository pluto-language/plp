package operators

import (
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

func getRoot() (string, error) {
	var root string

	if r, exists := os.LookupEnv("PLUTO"); exists {
		root = r
	} else {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		root = filepath.Join(usr.HomeDir, "pluto")
	}

	return root, nil
}

func downloadDirectory() error {
	var (
		url       = "https://raw.githubusercontent.com/pluto-language/packages/master/packages.json"
		root, err = getRoot()
	)

	if err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(root, "libraries", "packages.json"))
	if err != nil {
		return err
	}

	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

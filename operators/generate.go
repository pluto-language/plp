package operators

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	git "gopkg.in/src-d/go-git.v4"
)

type genOpts struct {
	title       string
	display     string
	description string
	author      string
	version     string
}

// Generate generates a package
// by cloning the package template repo
func Generate() error {
	opts, err := promptOptions()
	if err != nil {
		return err
	}

	pwd, err := filepath.Abs(".")
	if err != nil {
		return err
	}

	dir := filepath.Join(pwd, opts.title)

	fmt.Println()

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/pluto-language/package-template",
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	err = os.Remove(filepath.Join(dir, "LICENSE"))
	if err != nil {
		return err
	}

	err = applyOptionsToTemplate(dir, opts)
	if err != nil {
		return err
	}

	fmt.Printf("\nPackage '%s' has been created!\n", opts.title)

	return nil
}

func promptOptions() (*genOpts, error) {
	reader := bufio.NewReader(os.Stdin)
	colour := color.New(color.Bold)

	title, err := promptOption("package title", reader, colour)
	if err != nil {
		return nil, err
	}

	display, err := promptOption("display name", reader, colour)
	if err != nil {
		return nil, err
	}

	description, err := promptOption("description", reader, colour)
	if err != nil {
		return nil, err
	}

	author, err := promptOption("author", reader, colour)
	if err != nil {
		return nil, err
	}

	version, err := promptOption("version", reader, colour)
	if err != nil {
		return nil, err
	}

	return &genOpts{
		title:       title,
		display:     display,
		description: description,
		author:      author,
		version:     version,
	}, nil
}

func promptOption(msg string, reader *bufio.Reader, colour *color.Color) (string, error) {
	colour.Printf("%15s: ", msg)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}

func applyOptionsToTemplate(dir string, opts *genOpts) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == ".git" {
			continue
		}

		path := filepath.Join(dir, file.Name())

		if file.IsDir() {
			err = applyOptionsToTemplate(path, opts)
			if err != nil {
				return err
			}
		} else {
			var (
				name    = file.Name()
				newname = applyOptions(name, opts)
				newpath = filepath.Join(dir, newname)
			)

			err = os.Rename(path, newpath)
			if err != nil {
				return err
			}

			handle, err := os.Open(newpath)
			if err != nil {
				return err
			}

			data, err := ioutil.ReadAll(handle)
			if err != nil {
				return err
			}

			newstr := applyOptions(string(data), opts)

			err = ioutil.WriteFile(newpath, []byte(newstr), 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func applyOptions(str string, opts *genOpts) string {
	replacements := map[string]string{
		"%title":       opts.title,
		"%display":     opts.display,
		"%description": opts.description,
		"%author":      opts.author,
		"%version":     opts.version,
	}

	for temp, opt := range replacements {
		str = strings.Replace(str, temp, opt, -1)
	}

	return str
}

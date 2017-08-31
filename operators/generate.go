package operators

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

	// A title is entirely composed of lowercase letters, digits, dashes, and underscores
	title, err := promptOption(
		"package title",
		`^[a-z\d-_]+$`,
		"the title can only contain lowercase letters, digits, dashes, and underscores",
		reader,
		colour,
	)

	if err != nil {
		return nil, err
	}

	display, err := promptOption(
		"display name",
		`^.+$`,
		"the display name must be at least 1 character",
		reader,
		colour,
	)
	if err != nil {
		return nil, err
	}

	description, err := promptOption(
		"description",
		`^.+$`,
		"the description must be at least 1 character",
		reader,
		colour,
	)
	if err != nil {
		return nil, err
	}

	author, err := promptOption(
		"author",
		`^.+$`,
		"the author's name must be at least 1 character",
		reader,
		colour,
	)
	if err != nil {
		return nil, err
	}

	version, err := promptOption(
		"version",
		`^\d+\.\d+\.\d+$`,
		"the version must be three integers separated by dots",
		reader,
		colour,
	)
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

func promptOption(msg, pattern, errMsg string, reader *bufio.Reader, colour *color.Color) (string, error) {
	var (
		text string
		err  error
	)

	for {
		colour.Printf("%15s: ", msg)

		text, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		text = strings.TrimSpace(text)

		matches, err := regexp.MatchString(pattern, text)
		if err != nil {
			return "", err
		}

		if matches {
			goto done
		}

		color.Red("  invalid input: %s", errMsg)
	}

done:
	return text, nil
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

package operators

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
)

type pkg struct {
	title, repo string
}

type pkgs []pkg

func (p pkgs) Len() int           { return len(p) }
func (p pkgs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p pkgs) Less(i, j int) bool { return p[i].title < p[j].title }

// List lists all available packages
// from packages.json
func List() error {
	downloadDirectory()

	packages, err := getPackages()
	if err != nil {
		return err
	}

	longest := longestRepo(packages)

	headers := fmt.Sprintf("\n%-20s  %-"+fmt.Sprintf("%d", longest)+"s\n", "TITLE", "REPO")
	headers = strings.Repeat(" ", len(headers)-2) + headers

	headerColour := color.New(color.FgCyan, color.Underline, color.Bold)
	headerColour.Printf(headers)

	sorted := sortPackages(packages)

	for _, p := range sorted {
		fmt.Println(packageString(p))
	}

	return nil
}

func packageString(pkg pkg) string {
	return fmt.Sprintf("%-20s  %s", pkg.title, pkg.repo)
}

func longestRepo(pkgs map[string]string) int {
	longest := 0

	for _, repo := range pkgs {
		l := len(repo)

		if l > longest {
			longest = l
		}
	}

	return longest
}

func sortPackages(packages map[string]string) pkgs {
	var pks pkgs

	for title, repo := range packages {
		pks = append(pks, pkg{
			title: title,
			repo:  repo,
		})
	}

	sort.Sort(pks)
	return pks
}

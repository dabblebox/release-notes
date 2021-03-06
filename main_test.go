package main

import (
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/dabblebox/release-notes/components/links"
	"github.com/dabblebox/release-notes/components/notes"

	"github.com/dabblebox/release-notes/components/integration/git"

	"github.com/blang/semver"
)

func TestBuilder(t *testing.T) {

	changes, err := notes.Build("ams-guardian-api", "v0.2.0-beta", `[A-Z]{7}-\d*`, os.Getenv("GITHUB_URL"), os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"), 10, links.URL{
		ReplaceRegEx:  `[A-Z]{7}-\d*`,
		TokenizedLink: "<http://tickets.turner.com/browse/{TICKET_NUMBER}|{TICKET_NUMBER}>",
		ReplaceToken:  "TICKET_NUMBER",
	})
	if err != nil {
		fmt.Println(err)
	}

	summary := notes.Format(changes)

	fmt.Println(summary)
}

func TestGitHubGetTags(t *testing.T) {

	tags, err := git.GetTags("ams-guardian-api", os.Getenv("GITHUB_URL"), os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	commit, err := git.GetCommit("ams-guardian-api", tags["v0.2.0"].Commit.SHA, os.Getenv("GITHUB_URL"), os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tags)
	fmt.Println(commit)
}

func TestSemverSort(t *testing.T) {
	v1, err := semver.Make("1.0.0-beta.1234")
	if err != nil {
		fmt.Print(err)
	}

	v2, err := semver.Make("2.0.0-beta.1234")
	if err != nil {
		fmt.Print(err)
	}

	v3, err := semver.Make("3.0.0-beta.1234")
	if err != nil {
		fmt.Print(err)
	}

	versions := semver.Versions{
		v1,
		v3,
		v2,
	}

	sort.Sort(versions)

	fmt.Print(versions)
}

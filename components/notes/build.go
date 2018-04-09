package notes

import (
	"fmt"
	"regexp"

	"github.com/dabblebox/release-notes/components/integration/git"
	"github.com/dabblebox/release-notes/components/links"
)

// Build ...
func Build(repo, tag, filter, url, accessToken string, maxCommits int, linkConfig links.URL) ([]string, error) {

	tags, err := git.GetTags(repo, url, accessToken)
	if err != nil {
		return []string{}, err
	}

	startTag, found := tags[tag]
	if !found {
		return []string{}, fmt.Errorf("tag %s not found", tag)
	}

	commit, err := git.GetCommit(repo, startTag.Commit.SHA, url, accessToken)
	if err != nil {
		return []string{}, err
	}

	notes := []string{}
	for x := 1; x <= maxCommits; x++ {

		if isCommitImportant(filter, commit.Commit) {
			commit.Commit.Message = links.Insert(linkConfig, commit.Commit.Message)

			notes = append(notes, commit.Commit.Message)
		} else {
			x--
		}

		if isCommitTagged(commit.Parents[0].SHA, tags) {
			break
		}

		commit, err = git.GetCommit(repo, commit.Parents[0].SHA, url, accessToken)
		if err != nil {
			return notes, err
		}
	}

	return notes, nil
}

func isCommitImportant(filter string, commit git.CommitCommit) bool {
	re := regexp.MustCompile(filter)

	return re.MatchString(commit.Message)
}

func isCommitTagged(SHA string, tags map[string]git.Tag) bool {
	for _, t := range tags {
		if t.Commit.SHA == SHA {
			return true
		}
	}

	return false
}

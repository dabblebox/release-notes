package notes

import (
	"fmt"

	"github.com/dabblebox/release-notes/components/integration/git"
	"github.com/spf13/viper"
)

// Build ...
func Build(repo, tag string) ([]string, error) {
	tags, err := git.GetTags(repo)
	if err != nil {
		return []string{}, err
	}

	startTag, found := tags[tag]
	if !found {
		return []string{}, fmt.Errorf("tag %s not found", tag)
	}

	commit, err := git.GetCommit(repo, startTag.Commit.SHA)
	if err != nil {
		return []string{}, err
	}

	notes := []string{}
	for x := 1; x <= viper.GetInt("max-commits"); x++ {

		notes = append(notes, commit.Commit.Message)

		if isCommitTagged(commit.Parents[0].SHA, tags) {
			break
		}

		commit, err = git.GetCommit(repo, commit.Parents[0].SHA)
		if err != nil {
			return notes, err
		}
	}

	return notes, nil
}

func isCommitTagged(SHA string, tags map[string]git.Tag) bool {
	for _, t := range tags {
		if t.Commit.SHA == SHA {
			return true
		}
	}

	return false
}

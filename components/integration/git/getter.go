package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/spf13/viper"
)

// Tag ...
type Tag struct {
	Name   string
	Commit TagCommit
}

// TagCommit ...
type TagCommit struct {
	SHA string
}

// GetTags ...
func GetTags(repo string) (map[string]Tag, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repos/turnercode/%s/tags", viper.GetString("github-url"), repo), nil)
	if err != nil {
		return map[string]Tag{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", viper.GetString("github-auth")))
	resp, err := client.Do(req)
	if err != nil {
		return map[string]Tag{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return map[string]Tag{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]Tag{}, err
	}

	tags := []Tag{}
	if err := json.Unmarshal(body, &tags); err != nil {
		return map[string]Tag{}, err
	}

	tagMap := map[string]Tag{}
	for _, tag := range tags {
		tagMap[tag.Name] = tag
	}

	return tagMap, nil
}

// Commit ...
type Commit struct {
	SHA     string
	Parents []CommitParent
	Commit  CommitCommit
}

// CommitParent ...
type CommitParent struct {
	SHA string
}

// CommitCommit ...
type CommitCommit struct {
	Message string
}

// GetCommit ...
func GetCommit(repo, SHA string) (Commit, error) {
	client := &http.Client{}

	commit := Commit{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repos/turnercode/%s/commits/%s", viper.GetString("github-url"), repo, SHA), nil)
	if err != nil {
		return commit, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("token %s", viper.GetString("github-auth")))
	resp, err := client.Do(req)
	if err != nil {
		return commit, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return commit, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return commit, err
	}

	err = json.Unmarshal(body, &commit)

	return commit, err
}

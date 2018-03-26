// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/dabblebox/release-notes/components/links"
	"github.com/dabblebox/release-notes/components/notes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitTag  = ""
	gitRepo = ""

	githubURL    = ""
	githubAPIKey = ""

	maxCommits = 0
	filter     = ""

	urlRegEx = ""
	urlToken = ""
	urlLink  = ""
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate release notes from commits.",
	Long:  `Generate release notes from commits.`,
	Run: func(cmd *cobra.Command, args []string) {

		url := links.URL{
			ReplaceRegEx:  viper.GetString("url-regex"),
			ReplaceToken:  viper.GetString("url-token"),
			TokenizedLink: viper.GetString("url-link"),
		}

		changes, err := notes.Build(
			viper.GetString("git-repo"),
			viper.GetString("git-tag"),
			viper.GetString("commit-filter"),
			viper.GetInt("max-commits"),
			url)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		summary := notes.Format(changes)

		fmt.Fprintf(os.Stdout, summary)
	},
}

func init() {
	RootCmd.AddCommand(genCmd)

	const tagName = "git-tag"
	genCmd.PersistentFlags().StringVarP(&gitTag, tagName, "t", "", "git release tag")
	viper.BindPFlag(tagName, genCmd.PersistentFlags().Lookup(tagName))

	const repoName = "git-repo"
	genCmd.PersistentFlags().StringVarP(&gitRepo, repoName, "r", "", "git repo")
	viper.BindPFlag(repoName, genCmd.PersistentFlags().Lookup(repoName))

	const urlName = "github-url"
	genCmd.PersistentFlags().StringVarP(&githubURL, urlName, "u", "https://api.github.com", "GitHub api url")
	viper.BindPFlag(urlName, genCmd.PersistentFlags().Lookup(urlName))

	const authName = "github-auth"
	genCmd.PersistentFlags().StringVarP(&githubAPIKey, authName, "a", "", "GitHub api authorization token")
	viper.BindPFlag(authName, genCmd.PersistentFlags().Lookup(authName))

	const maxCommitsName = "max-commits"
	genCmd.PersistentFlags().IntVarP(&maxCommits, maxCommitsName, "c", 100, "max number of commits to display")
	viper.BindPFlag(maxCommitsName, genCmd.PersistentFlags().Lookup(maxCommitsName))

	const commitFilterName = "commit-filter"
	genCmd.PersistentFlags().StringVarP(&filter, commitFilterName, "f", "", "regex filter that removes commits that do not match")
	viper.BindPFlag(commitFilterName, genCmd.PersistentFlags().Lookup(commitFilterName))

	const urlRegexName = "url-regex"
	genCmd.PersistentFlags().StringVarP(&urlRegEx, urlRegexName, "x", "", "regular expression for replacing token in link")
	viper.BindPFlag(urlRegexName, genCmd.PersistentFlags().Lookup(urlRegexName))

	const urlTokenName = "url-token"
	genCmd.PersistentFlags().StringVarP(&urlToken, urlTokenName, "k", "", "token in the link that will be replaced using regex")
	viper.BindPFlag(urlTokenName, genCmd.PersistentFlags().Lookup(urlTokenName))

	const urlLinkName = "url-link"
	genCmd.PersistentFlags().StringVarP(&urlLink, urlLinkName, "l", "", "link to use in the commit")
	viper.BindPFlag(urlLinkName, genCmd.PersistentFlags().Lookup(urlLinkName))
}

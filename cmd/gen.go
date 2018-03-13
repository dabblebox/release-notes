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

	"github.com/dabblebox/release-notes/components/notes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tag = ""
var repo = ""
var githubURL = ""
var githubAPIKey = ""
var maxCommits = 0

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate release notes from commits.",
	Long:  `Generate release notes from commits.`,
	Run: func(cmd *cobra.Command, args []string) {
		changes, err := notes.Build(viper.GetString("git-repo"), viper.GetString("git-tag"))
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
	genCmd.PersistentFlags().StringVarP(&tag, tagName, "t", "", "release tag")
	viper.BindPFlag(tagName, genCmd.PersistentFlags().Lookup(tagName))

	const repoName = "git-repo"
	genCmd.PersistentFlags().StringVarP(&repo, repoName, "r", "", "repo")
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
}

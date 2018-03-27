# Release Notes Generator #

This tool generates and writes GitHub commit release notes preceeding and including a specific tag to `stdout` where they can be fed into any continuous integration pipeline.

Remember to build targeting the supported OS and platform of your CI server.

### Examples ###

#### Build and Run ####

```bash 
$ go build
$ ./release-notes gen -t {RELEASE_TAG} -r {REPO_NAME} -a {GITHUB_PERSONAL_ACCESS_TOKEN} -f "[A-Z]{7}-\d*"
```

#### Build for CircleCI Linux ####

```bash 
$ env GOOS=linux GOARCH=386 go build
```

#### Send Release Notes to Slack ####

1. Pleace release notes binary in tools folder within your project.
2. Set up Slack [webhook](https://taylorcoding.slack.com/apps/new/A0F7XDUAZ-incoming-webhooks).
3. Add environment variables.
4. Call command from deploy scripts.
```bash 
$ echo '{"text":"Version '$VERSION_NUM' for '$CIRCLE_PROJECT_REPONAME' has been deployed to dev. '$(./tools/release-notes gen --git-repo $CIRCLE_PROJECT_REPONAME --git-tag $CIRCLE_TAG --github-auth $GIT_HUB_PERSONAL_ACCESS_TOKEN --url-link "<http://tickets.turner.com/browse/{TICKET_NUMBER}|{TICKET_NUMBER}>" --url-regex "[A-Z]{7}-\d*" --commit-filter "[A-Z]{7}-\d*" --url-token "{TICKET_NUMBER}")'"}' | curl -H "Content-Type:application/json" -d @- $SLACK_WEBHOOK_URL
```

### CLI Flags ###

| Short | Long | Default | Description |
|------|-------|-------|-------------|
|-t|--git-tag||git release tag including desired commits|
|-r|--git-repo||git repository including desired commits|
|-u|--github-url|https://api.github.com|GitHub api url|
|-a|--github-auth||GitHub personal access token|
|-c|--max-commits|25|number of commits to walk through searching for a previous release tag|
|-f|--commit-filter||regex filter to include specific commits|
|-x|--url-regex||regex used to search and replace url token in story link|
|-k|--url-token||url token in the story link that will be replaced using regex|
|-l|--url-link||story link injected into the commit|

### Supported Git Hosted Solutions ###

- GitHub
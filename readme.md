# Release Notes Generator #

This tool generates and writes GitHub commit release notes preceeding and including a specific tag to `stdout` where they can be fed into any continuous integration pipeline.

Remember to build targeting the supported OS and platform of your CI server.

### Example ###

#### Build and Run ####

```bash 
$ go build
$ ./release-notes gen -t {RELEASE_TAG} -r {REPO_NAME} -a {GITHUB_PERSONAL_ACCESS_TOKEN} -f "[A-Z]{7}-\d*"
```

#### Build for CircleCI Linux ####

```bash 
$ env GOOS=linux GOARCH=386 go build
```

### CLI Flag Definitions ###

| Short | Long | Default | Description |
|------|-------|-------|-------------|
|-t|--git-tag||git release tag|
|-r|--git-repo||git repo|
|-u|--github-url|https://api.github.com|GitHub api url|
|-a|--github-auth||GitHub api authorization token|
|-c|--max-commits|100|max number of commits to display|
|-f|--commit-filter||regex filter that removes commits that do not match|
|-x|--url-regex||regular expression for replacing token in link|
|-k|--url-token||token in the link that will be replaced using regex|
|-l|--url-link||link to use in the commit|

### Supported Git Solutions ###

- GitHub
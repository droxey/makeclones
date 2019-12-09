# makeclones

[![Go Report Card](https://goreportcard.com/badge/github.com/droxey/makeclones)](https://goreportcard.com/report/github.com/droxey/makeclones) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/7ed40f9f3ecf46709879d5fbac28fd9b)](https://www.codacy.com/app/droxey/makeclones?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=droxey/makeclones&amp;utm_campaign=Badge_Grade)

`git clone` repositories (in `username/reponame` format) from a google sheet column!

### Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)

## Installation

### Homebrew (Recommended)

```bash
brew tap droxey/makeclones
brew install droxey/makeclones/makeclones
```

### From Source

```bash
$ go build && go install
```

## Usage

```bash
$ makeclones
  -column string
        Column to scrape. Make sure data is in the format username/reponame (Required)
  -sheet string
        Google Sheets spreadsheet ID (Required)
  -skip int
        Skip a number of rows to accomodate headers
  -tab int
        Spreadsheet tab to look for the specified column
  -token string
        GitHub Personal Access Token (Create one at https://github.com/settings/tokens/new) with full control of private repositories (Required)
```

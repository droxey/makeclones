package main

import (
	"flag"
	"os"
)

func main() {
	// start := time.Now()

	var sheet string
	var column string
	var authToken string
	var skipRows int
	var tabIndex int
	var runSonar bool

	flag.StringVar(&sheet, "sheet", "", "Google Sheets spreadsheet ID (Required)")
	flag.StringVar(&column, "column", "", "Column to scrape. Make sure data is in the format username/reponame (Required)")
	flag.StringVar(&authToken, "token", "", "GitHub Personal Access Token (Create one at https://github.com/settings/tokens/new) with full control of private repositories (Required)")
	flag.IntVar(&skipRows, "skip", 0, "Skip a number of rows to accomodate headers")
	flag.IntVar(&tabIndex, "tab", 0, "Spreadsheet tab to look for the specified column")
	flag.BoolVar(&runSonar, "analyze", false, "EXPERIMENTAL. Add to SonarQube for analysis.")

	flag.Parse()

	if sheet == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if column == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if authToken == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	MakeClones(sheet, tabIndex, column, authToken, skipRows, runSonar)
	// display the number of repos and length of time in the terminal
	// 	since := time.Since(start).Seconds()
	// 	info("Cloned %d repos in %2f seconds", numOfReposCloned, since)
}

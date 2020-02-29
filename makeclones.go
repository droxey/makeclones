package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	s "gopkg.in/Iwark/spreadsheet.v2"
	g "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type repoDirPair struct {
	RepoURL string
	Dir     string
	Token   string
}

// var mux sync.Mutex

// initialize variable to store the number of repos cloned off of the initial sheet.
// for benchmarking purposes
var numOfReposCloned int

// MakeClones from a spreadsheet column
// Initial Benchmark: 21 reopos in 246 seconds
func MakeClones(sheetID string, tabIndex int, column string, token string, skip int, analyze bool) {
	service, err := s.NewService()
	checkIfError(err)

	sheets, err := service.FetchSpreadsheet(sheetID)
	checkIfError(err)

	sheet, err := sheets.SheetByIndex(uint(tabIndex))
	checkIfError(err)

	repoDirChan := make(chan repoDirPair, 50)
	results := make(chan bool, 50)

	var numOfDirs int

	for i := 1; i <= 9; i++ {
		go cloneWorker(repoDirChan, results)
	}

	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Row > uint(skip) {
				cellPos := cell.Pos()
				if string(cellPos[0]) == column && len(cell.Value) > 0 {
					checkIfError(err)

					prefix := "github.com/"
					directory := "github.com/" + cell.Value
					repoURL := cell.Value

					if !strings.HasPrefix(repoURL, "https://"+prefix) {
						repoURL = "https://" + prefix + cell.Value
					}

					warning("creating directory %s...", directory)
					err = os.MkdirAll(directory, os.ModePerm)
					checkIfError(err)
					numOfDirs++

					repoDirChan <- repoDirPair{Dir: directory, RepoURL: repoURL, Token: token}

				}
			}
		}
	}
	close(repoDirChan)
	for a := 1; a <= numOfDirs; a++ {
		<-results
	}
}

func cloneWorker(pairs chan repoDirPair, results chan<- bool) {
	for pair := range pairs {
		results <- clone(pair.Token, pair.RepoURL, pair.Dir)
	}
}

func clone(token, repoURL, directory string) bool {
	info("cloning %s into %s...", repoURL, directory)
	_, err := g.PlainClone(directory, false, &g.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "makeclones", // This can be anything except an empty string
			Password: token,
		},
		URL:      repoURL,
		Progress: os.Stdout,
	})
	checkIfError(err)
	numOfReposCloned++
	return true
}

// checkIfError should be used to naively panic if an error is not nil.
func checkIfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("[makeclones] \x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("%s", err))
	//os.Exit(1)
}

// info should be used to describe the example commands that are about to run.
func info(format string, args ...interface{}) {
	fmt.Printf("[makeclones] \x1b[32;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// warning should be used to display a warning
func warning(format string, args ...interface{}) {
	fmt.Printf("[makeclones] \x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

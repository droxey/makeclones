package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	s "gopkg.in/Iwark/spreadsheet.v2"
	g "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// MakeClones from a spreadsheet column
// Initial Benchmark: 21 reopos in 246 seconds
func MakeClones(sheetID string, tabIndex int, column string, token string, skip int, analyze bool) {
	service, err := s.NewService()
	checkIfError(err)

	sheets, err := service.FetchSpreadsheet(sheetID)
	checkIfError(err)

	sheet, err := sheets.SheetByIndex(uint(tabIndex))
	checkIfError(err)

	// var reposToAnalyze []string

	// initialize variable to store the number of repos cloned off of the initial sheet.
	// for benchmarking purposes
	var numOfReposCloned int
	start := time.Now()

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
					// if analyze {
					// 	name := strings.Split(repoURL, "https://github.com/")[1]
					// 	reposToAnalyze = append(reposToAnalyze, name)
					// }
				}
			}
		}

	}
	// display the number of repos and length of time in the terminal
	since := time.Since(start).Seconds()
	info("Cloned %d repos in %2f seconds", numOfReposCloned, since)

	// if analyze {
	// 	// Run analysis syncronously.
	// 	// sonar-scanner does not support concurrent operations.
	// 	wg := new(sync.WaitGroup)
	// 	wg.Add(len(reposToAnalyze))

	// 	for _, repo := range reposToAnalyze {
	// 		analyzeCode(repo, wg)
	// 	}

	// 	wg.Wait()
	// }
}

// func analyzeCode(name string, wg *sync.WaitGroup) {
// 	sonarURL := os.Getenv("SONARQUBE_URL")
// 	sonarUser := os.Getenv("SONARQUBE_USERNAME")
// 	sonarPass := os.Getenv("SONARQUBE_PASSWORD")

// 	// // Connect to SonarQube instance.
// 	// client, err := sonargo.NewClient(sonarURL+"/api", sonarUser, sonarPass)
// 	// checkIfError(err)

// 	// // Create a new Sonarqube Project.
// 	// opts := &sonargo.ProjectsCreateOption{Branch: "master", Name: name, Project: name, Visibility: ""}
// 	// _, _, err = client.Projects.Create(opts)
// 	// checkIfError(err)

// 	// Run sonar-scanner in the project directory to initialize the scan.
// 	info("starting analysis: %s", name)
// 	cmd := exec.Command("sonar-scanner",
// 		"-Dsonar.projectKey="+name,
// 		"-Dsonar.sources=.",
// 		"-Dsonar.host.url="+sonarURL,
// 		"-Dsonar.login="+sonarUser,
// 		"-Dsonar.password="+sonarPass)
// 	cmd.Start()
// 	cmd.Wait()

// 	info("analysis complete: %s", name)
// 	wg.Done()
// }

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

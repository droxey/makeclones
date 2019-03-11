package main

import (
	"fmt"
	"os"

	s "gopkg.in/Iwark/spreadsheet.v2"
	g "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// MakeClones from a spreadsheet column
func MakeClones(sheetID string, tabIndex int, column string, token string, skip int) {
	service, err := s.NewService()
	checkIfError(err)

	sheets, err := service.FetchSpreadsheet(sheetID)
	checkIfError(err)

	sheet, err := sheets.SheetByIndex(uint(tabIndex))
	checkIfError(err)

	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Row > uint(skip) {
				cellPos := cell.Pos()
				if string(cellPos[0]) == column && len(cell.Value) > 0 {
					checkIfError(err)

					directory := "github.com/" + cell.Value
					repoURL := "https://github.com/" + cell.Value

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
				}
			}
		}
	}
}

// checkIfError should be used to naively panics if an error is not nil.
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
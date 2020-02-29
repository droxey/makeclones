package main

import (
	"os"
	"testing"
)

const (
	tabIndex = 0
	skip     = 0
	column   = "B"
	analyze  = false
)

func BenchmarkMakeClones(b *testing.B) {
	for i := 0; i < 1; i++ {
		MakeClones(os.Getenv("SHEETID"), tabIndex, column, os.Getenv("POATOKEN"), skip, analyze)
	}
}

package main

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	files := listFiles(".")
	if len(files) == 0 {
		t.Errorf("Files has nothing")
	}
}

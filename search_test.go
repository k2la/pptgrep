package main

import (
	"path/filepath"
	"testing"
)

func TestIsPpt(t *testing.T) {
	if !isPpt("hoge.pptx") {
		t.Errorf("hoge.pptx is PPT file.")
	}
}

func TestContainWord(t *testing.T) {
	archive, _ := filepath.Abs("./test.pptx")
	if !containWord(archive, "GOOGLE") {
		t.Errorf("test.pptx has google")
	}
}

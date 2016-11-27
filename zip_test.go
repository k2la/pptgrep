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

func TestUnzip(t *testing.T) {
	archive, _ := filepath.Abs("./test.pptx")
	target, _ := filepath.Abs("./test")
	err := unzip(archive, target)
	if err != nil {
		t.Errorf("Fail unzip pptx")
	}
}

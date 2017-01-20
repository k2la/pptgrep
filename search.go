package main

import (
	"archive/zip"
	"path/filepath"
	"regexp"
)

var slide = regexp.MustCompile("ppt/slides/slide")

func isPpt(path string) bool {
	e := filepath.Ext(path)
	if e == ".ppt" || e == ".pptx" {
		return true
	}
	return false
}

func containWord(archive, word string) bool {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return false
	}

	for _, file := range reader.File {
		if slide.MatchString(file.Name) {
			fileReader, err := file.Open()
			if err != nil {
				return false
			}
			var p = make([]byte, file.FileInfo().Size())
			fileReader.Read(p)
			defer fileReader.Close()
			r := regexp.MustCompile(word)
			if r.MatchString(string(p)) {
				return true
			}
		}
	}
	return false
}

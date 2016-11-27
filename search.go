package main

import (
	"archive/zip"
	"path/filepath"
	"regexp"
)

func isPpt(path string) bool {
	e := filepath.Ext(path)
	if e == ".ppt" || e == ".pptx" {
		return true
	}
	return false
}

func isContain(archive, word string) bool {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return false
	}

	for _, file := range reader.File {
		if file.Name == "docProps/app.xml" {
			fileReader, err := file.Open()
			if err != nil {
				return false
			}
			p := []byte{}
			for i := int64(0); i < file.FileInfo().Size(); i++ {
				p = append(p, 0)
			}
			fileReader.Read(p)
			defer fileReader.Close()
			r := regexp.MustCompile(word)
			return r.MatchString(string(p))
		}
	}
	return false
}

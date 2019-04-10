package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// FoundFile fill if found searched text in file
type FoundFile struct {
	Name  string
	Path  string
	Lines []FoundText
}

// FoundText fill if text is found
type FoundText struct {
	Line int
	Text string
}

// FindInFolder find text in folder files
func FindInFolder(find, folder string) ([]FoundFile, int) {

	var filesNum int
	var report []FoundFile

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && path != folder {
			filesNum++

			lines, _ := FindInFile(find, path)
			if len(lines) != 0 {

				found := FoundFile{
					Name:  info.Name(),
					Path:  path,
					Lines: lines,
				}

				report = append(report, found)

			}

		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return report, filesNum
}

// FindInFile open file, check for find text
func FindInFile(find, file string) ([]FoundText, error) {

	var lines []FoundText

	f, err := os.Open(file)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {

		fileText := scanner.Text()

		if strings.Contains(fileText, find) {

			lineText := FoundText{
				Line: line,
				Text: fileText,
			}

			lines = append(lines, lineText)
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
	}

	return lines, nil

}

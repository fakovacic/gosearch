package main

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Equanox/gotron"
)

// Search handle search request
type Search struct {
	Text   string
	Folder string
}

// CustomEvent Create a custom event struct that has a pointer to gotron.Event
type CustomEvent struct {
	*gotron.Event
	Msg string "json:msg"
}

// Results contain search results
type Results struct {
	*gotron.Event
	Results []FoundFile "json:results"
}

func main() {
	// Create a new browser window instance
	window, err := gotron.New("app")
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 800
	window.WindowOptions.Height = 600
	window.WindowOptions.Title = "GoSearch"

	onEvent := gotron.Event{Event: "search"}

	window.On(&onEvent, func(bin []byte) {

		var searchReq Search

		json.Unmarshal(bin, &searchReq)

		if searchReq.Text != "" && searchReq.Folder != "" {

			if _, err := os.Stat(searchReq.Folder); os.IsNotExist(err) {
				window.Send(&CustomEvent{
					Event: &gotron.Event{Event: "error"},
					Msg:   "Not found folder: " + searchReq.Folder,
				})
			}

			if _, err := os.Stat(searchReq.Folder); !os.IsNotExist(err) {
				start := time.Now()
				reportFiles, filesNum := FindInFolder(searchReq.Text, searchReq.Folder)

				elapsed := time.Since(start)

				reportCount := len(reportFiles)

				if reportCount != 0 {

					window.Send(&CustomEvent{
						Event: &gotron.Event{Event: "success"},
						Msg:   "Found " + searchReq.Text + " in " + strconv.Itoa(reportCount) + "/" + strconv.Itoa(filesNum) + " files, took " + elapsed.String(),
					})

					window.Send(&Results{
						Event:   &gotron.Event{Event: "results"},
						Results: reportFiles,
					})

				} else {

					window.Send(&CustomEvent{
						Event: &gotron.Event{Event: "error"},
						Msg:   "Not found in " + strconv.Itoa(filesNum) + ", took " + elapsed.String(),
					})

				}
			}

		} else {

			window.Send(&CustomEvent{
				Event: &gotron.Event{Event: "error"},
				Msg:   "Please fill required fields",
			})

		}

	})

	// Start the browser window.
	// This will establish a golang <=> nodejs bridge using websockets,
	// to control ElectronBrowserWindow with our window object.
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Open dev tools must be used after window.Start
	window.OpenDevTools()

	// Wait for the application to close
	<-done
}

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

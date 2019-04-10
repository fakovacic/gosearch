package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ArrayToString convert int slice to string
func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// TxtReport sum up find results, return report string
func TxtReport(reportFiles []FoundFile, text, folder string, filesNum int) string {

	var reportResults string

	reportCount := len(reportFiles)

	if reportCount != 0 {

		reportResults = reportResults + "Found " + text + " in " + strconv.Itoa(reportCount) + "/" + strconv.Itoa(filesNum) + " files"

		for _, ffile := range reportFiles {

			reportResults = reportResults + `

` + ffile.Name + ` - ` + ffile.Path + `
Lines:`

			for _, ln := range ffile.Lines {

				reportResults = reportResults + `
` + strconv.Itoa(ln.Line) + `: ` + ln.Text

			}

		}

	} else {

		reportResults = "Not found " + text + " in folder " + folder

	}

	return reportResults
}

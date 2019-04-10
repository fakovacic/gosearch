package main

import (
	"fmt"
	"os"
)

// WriteTxt create file, write string report from ReportResults
func WriteTxt(report, fileName string) {

	logFile, _ := os.Create(fileName + ".log")
	defer logFile.Close()

	logFile.WriteString(report)

	fmt.Println("Search report:" + fileName)

}

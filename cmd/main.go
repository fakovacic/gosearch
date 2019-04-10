package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// Define vars for search report
	text, folder, logger := Arger()
	// Find all text in folder files
	reportFiles, filesNum := FindInFolder(text, folder)
	// Sum up found files in report string
	txtResult := TxtReport(reportFiles, text, folder, filesNum)

	switch logger {
	case "txt":

		fileName := "gosearch_" + time.Now().Format("20060102150405")

		WriteTxt(txtResult, fileName)
		break
	}

	fmt.Println(txtResult)

	elapsed := time.Since(start)
	fmt.Println("Took ", elapsed)

}

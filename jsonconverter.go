package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func convertToJSON(f *os.File, isTSVFile bool) {
	rows := make([]map[string]string, 0)
	var columns = []string{"language", "frequency"}
	// Create a new reader.
	csvReader := csv.NewReader(bufio.NewReader(f))
	csvReader.TrimLeadingSpace = true
	if isTSVFile {
		csvReader.Comma = '\t' // Use tab-delimited instead of comma
	}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		row := make(map[string]string)
		for i, n := range columns {
			row[n] = record[i]
		}
		rows = append(rows, row)
	}
	data, err := json.MarshalIndent(&rows, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	// print the reformatted struct as JSON
	fmt.Printf("%s\n", data)
}

func main() {
	csvPtr := flag.String("csv-file", "", "CSV File to parse. (Required)")
	tsvPtr := flag.String("tsv-file", "", "TSV File to parse. (Required)")
	flag.Parse()

	if *csvPtr == "" && *tsvPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	} else if *csvPtr != "" {
		f, err := os.Open(*csvPtr)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		convertToJSON(f, false)
	} else {
		f, err := os.Open(*tsvPtr)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		convertToJSON(f, true)
	}
}

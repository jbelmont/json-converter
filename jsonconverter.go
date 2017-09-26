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

// FieldsReader reads fields from tab separated file
type FieldsReader struct {
	*csv.Reader
	fields []int
}

func (r *FieldsReader) Read() (record []string, err error) {
	rec, err := r.Reader.Read()
	if err != nil {
		return nil, err
	}

	record = make([]string, len(r.fields))
	for i, f := range r.fields {
		record[i] = rec[f]
	}

	return record, nil
}

func main() {
	rows := make([]map[string]string, 0)
	var columns = []string{"name", "email", "phone"}
	csvPtr := flag.String("csv-file", "", "CSV File to parse. (Required)")
	tsvPtr := flag.String("tsv-file", "", "CSV File to parse. (Required)")
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
		// Create a new reader.
		csvReader := csv.NewReader(bufio.NewReader(f))
		csvReader.TrimLeadingSpace = true
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
		fmt.Println(data)
	} else {
		f, err := os.Open(*tsvPtr)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		reader := csv.NewReader(bufio.NewReader(f))
		reader.Comma = '\t' // Use tab-delimited instead of comma <---- here!
		tsvFile, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for field, each := range tsvFile {
			fmt.Println(field, each)
		}
	}
}

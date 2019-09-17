package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	reader := csv.NewReader(os.Stdin)

	var headers []string
	var data []map[string]string

	for i := 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read csv: %s", err)
		}

		if i == 0 {
			// first line. gather headers
			headers = record
		} else {
			vals := map[string]string{}
			for j := 0; j < len(headers); j++ {
				vals[headers[j]] = record[j]
			}
			data = append(data, vals)
		}
	}

	output, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("can't unmarshal data: %s", err)
	}

	fmt.Println(string(output))
}

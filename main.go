package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

func main() {
	filename := os.Getenv("input")
	key := os.Getenv("key")
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	if len(records) < 1 {
		log.Fatal("The csv file should have header")
	}

	header := records[0]
	cared_header := make(map[string]int)
	for i, field := range header {
		cared_header[field] = i
	}

	delete(cared_header, "_")

	if _, ok := cared_header[key]; !ok {
		log.Fatal("There is not field as key")
	}

	contents := records[1:]

	w := json.NewEncoder(os.Stdout)
	for _, record := range contents {
		data := make(map[string]interface{})
		for field, i := range cared_header {
			data[field] = record[i]
			if field == key {
				data["_key"] = record[i]
			}
		}
		w.Encode(data)
	}
}

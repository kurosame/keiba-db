package service

import (
	"encoding/csv"
	"log"
	"os"
)

// OutputCSV is output CSV to dataset directory
func OutputCSV(row []string) {
	file, err := os.Create("./dataset/netkeiba.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write(row)
	writer.Flush()
}

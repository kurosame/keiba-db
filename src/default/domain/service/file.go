package service

import (
	"bufio"
	"log"
	"os"
)

// OutputJSONL is output jsonl to dataset directory
func OutputJSONL(rows []string) {
	file, err := os.Create("./dataset/netkeiba.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, r := range rows {
		_, err := writer.WriteString(r + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()
}

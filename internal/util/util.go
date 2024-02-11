package util

import (
	"bufio"
	"log"
	"os"
)

// SaveToFile saves the string to a file.
func SaveToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()
}

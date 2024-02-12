package util

import (
	"bufio"
	"log/slog"
	"os"
)

// SaveToFile saves the string to a file.
func SaveToFile(filename string, data string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		slog.Error("Error creating file:", err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString(data)
	if err != nil {
		slog.Error("Error writing to file:", err)
		return
	}
	w.Flush()
	return nil
}

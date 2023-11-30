package internal

import (
	"bufio"
	"os"
)

// ReadFile takes a filepath as an argument and returns a string
// slice of the contents of the file
func ReadFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	output := make([]string, 0)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output, err
}

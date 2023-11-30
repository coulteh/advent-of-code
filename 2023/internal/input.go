package internal

import "path/filepath"

const inputPath = "2023/input/"

func GetInputPath(s string) (string, error) {
	base, err := filepath.Abs(inputPath)
	if err != nil {
		return "", err
	}
	return filepath.Join(base, s), nil
}

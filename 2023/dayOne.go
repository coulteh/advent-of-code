package twenty23

import (
	"advent-of-code/2023/internal"
	"io"
	"strconv"
	"unicode"
)

var inputFile, _ = internal.GetInputPath("dayOne.txt")

func GetInput() (io.Reader, error) {
	return internal.ReadFile(inputFile)
}

func DayOne(r io.Reader, part2 bool) (output string, err error) {
	input := internal.ReaderToStrings(r)
	if err != nil {
		return "", err
	}
	var total int
	for _, line := range input {
		ints := make([]int, 0)
		for _, char := range line {
			if unicode.IsDigit(char) {
				ints = append(ints, int(char-'0'))
			}
		}
		output = strconv.Itoa(ints[0]) + strconv.Itoa(ints[len(ints)-1])
		outputInt, err := strconv.Atoi(output)
		if err != nil {
			return "", err
		}
		total += outputInt
		output = strconv.Itoa(total)
	}
	//output = strconv.Itoa(total)
	if part2 {
	}
	return output, nil
}

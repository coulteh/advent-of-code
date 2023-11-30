package twenty23

import (
	"advent-of-code/2023/internal"
	"fmt"
)

var inputFile, _ = internal.GetInputPath("dayOne.txt")

func DayOne(part2 bool) error {
	input, err := internal.ReadFile(inputFile)
	if err != nil {
		return err
	}
	for _, line := range input {
		fmt.Println(line)
	}
	if part2 {
		fmt.Println("part 2")
	}
	return nil
}

package twenty23

import (
	"errors"
)

func DaySelector(day int, part2 bool) (output string, err error) {
	if day < 1 || day > 24 {
		return output, errors.New("day must be in range 1-24")
	}
	switch day {
	case 1:
		r, err := GetInput()
		if err != nil {
			return "", err
		}
		output, err = DayOne(r, part2)
	default:
		err = errors.New("day not implemented yet")
	}
	return output, err
}

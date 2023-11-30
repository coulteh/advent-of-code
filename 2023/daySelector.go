package twenty23

import (
	"errors"
)

func DaySelector(day int, part2 bool) (err error) {
	if day < 1 || day > 24 {
		return errors.New("day must be in range 1-24")
	}
	switch day {
	case 1:
		err = DayOne(part2)
	default:
		err = errors.New("day not implemented yet")
	}
	return
}

# Advent Of Code

This project contains my Advent of Code entries written in Go. At time of writing, this only contains my 2023 entry. The basis is a [Cobra](https://cobra.dev/) application to generate a given day's challenges.

## Usage

The primary command is invoked with `advent <year> <day>`. There is one additional flag available:
* `--part2` - Runs the second part of the given day

Attempting to invoke an unimplemented day will result in an error.

## Extending

Adding an additional day is relatively straightforward. Some boilerplate is required:
```
func DayX(part2 bool) error {
	input, err := internal.ReadFile(inputFile)
	if err != nil {
		return err
	}
	...Part one logic goes here...
	if part2 {
		...Part two logic goes here...
	}
```
And `daySelector.go` needs the following added to the `switch` in `DaySelector()`:
```
switch day {
	case 1:
		err = DayOne(part2)
	case X:
	    err = DayX(part2)
	default:
		err = errors.New("day not implemented yet")
	}
```
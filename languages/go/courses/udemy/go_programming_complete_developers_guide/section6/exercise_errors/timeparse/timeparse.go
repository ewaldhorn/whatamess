package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid time input", input}
	} else {
		hours, err := strconv.Atoi(components[0])

		if err != nil {
			return Time{}, &TimeParseError{"Unable to parse hours", input}
		}

		minutes, err := strconv.Atoi(components[1])

		if err != nil {
			return Time{}, &TimeParseError{"Unable to parse minutes", input}
		}

		seconds, err := strconv.Atoi(components[2])

		if err != nil {
			return Time{}, &TimeParseError{"Unable to parse seconds", input}
		}

		if hours < 0 || hours > 23 {
			return Time{}, &TimeParseError{"Invalid hour specified", input}
		}
		if minutes < 0 || minutes > 59 {
			return Time{}, &TimeParseError{"Invalid minutes specified", input}
		}
		if seconds < 0 || seconds > 59 {
			return Time{}, &TimeParseError{"Invalid seconds specified", input}
		}

		return Time{hours, minutes, seconds}, nil
	}
}

package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	switch {
	case number%3 == 0:
		return "Pling"
	case number%5 == 0:
		return "Plang"
	case number%7 == 0:
		return "Plong"
	default:
		return strconv.Itoa(number)
	}
}

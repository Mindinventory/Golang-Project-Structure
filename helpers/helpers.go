package helpers

import (
	"strconv"
)

//Any helper functions

func Int64ToString(inputNum int64) string {
	// to convert a float number to a string
	return strconv.FormatInt(inputNum, 10)
}

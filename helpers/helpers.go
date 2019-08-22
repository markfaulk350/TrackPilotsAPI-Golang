package helpers

import "strconv"

func Int64ToString(num int64) string {
	s := strconv.FormatInt(num, 10)
	return s
}

func Float64ToString(num float64) string {
	s := strconv.FormatFloat(num, 'f', -1, 64)
	return s
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

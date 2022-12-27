package util

import (
	"strconv"
	"time"
)

func GetYear() int {
	return time.Now().Year()
}

func GetYearString() string {
	return strconv.Itoa(GetYear())
}

func GetMonth() int {
	return int(time.Now().Month())
}

func GetMonthString() string {
	return strconv.Itoa(GetMonth())
}

func GetDay() int {
	return time.Now().Day()
}

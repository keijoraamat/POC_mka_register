package helpers

import (
	"strconv"
	"strings"
	"time"
)

func ParseDate(dateStr string) (date time.Time) {

	if dateStr == "dd.mm.yyyy" {
		dateStr = time.Now().Format("02.01.2006")
	}

	dateParts := splitByDot(dateStr)

	day, month, year := convertToObj(dateParts)

	date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return
}

func convertToObj(dateParts []string) (day, month, year int) {
	day, _ = strconv.Atoi(dateParts[0])
	month, _ = strconv.Atoi(dateParts[1])
	year, _ = strconv.Atoi(dateParts[2])

	return
}

func splitByDot(dateStr string) []string {
	var delimiter = "."
	return strings.Split(dateStr, delimiter)
}

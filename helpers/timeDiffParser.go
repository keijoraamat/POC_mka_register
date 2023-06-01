package helpers

import (
	"math"
	"time"
)

func WeeksToEnd(startTime time.Time) (weeks float64) {

	var endTime = startTime.AddDate(0, 6, 0)

	days := int(time.Until(endTime).Hours() / 24)
	weeks = roundFloat(float64(days/7), 2)

	return

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

package helpers

import (
	"log"
	"strconv"
)

func ParseBool(strBool string) (boolVal bool) {

	if len(strBool) > 0 {
		boolVal, err := strconv.ParseBool(strBool)
		if err != nil {
			log.Fatal("Could not convert string to bool ", err)
		}
		return boolVal
	}

	boolVal = false
	return
}

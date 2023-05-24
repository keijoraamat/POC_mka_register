package helpers

func ParseCheckBox(checkBox string) (boolVal bool) {

	if checkBox == `on` {
		boolVal = true
	}

	boolVal = false
	return
}

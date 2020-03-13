package bin

func AddUnderline(str string, c string) (result string) {
	result = str + "\n"

	for i := 0; i < len(str); i++ {
		result = result + c
	}

	return result + "\n"
}

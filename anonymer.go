package anonymer

func Anonymize(name string) string {
	if name == "" {
		return ""
	}

	newBytes := make([]rune, 5)
	for i := range newBytes {
		newBytes[i] = '*'
	}

	first := name[0]
	last := name[len(name)-1]
	newBytes[0] = rune(first)
	newBytes[4] = rune(last)

	return string(newBytes)
}

package helpers

func RemovePrefix(str string) string {
	strRune := []rune(str)
	return string(strRune[2:])
}

func RemovePrefixFromAddress(str string) string {
	strRune := []rune(str)
	return string(strRune[5:])
}

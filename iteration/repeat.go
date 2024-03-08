package iteration

const repeatCount = 5

func Repeat(character string) string {
	ret := ""
	for range repeatCount {
		ret += character
	}
	return ret
}

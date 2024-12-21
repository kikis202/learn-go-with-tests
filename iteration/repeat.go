package iteration

func Repeat(character string, repeatCount int) (repeated string) {
	for range repeatCount {
		repeated += character
	}
	return repeated
}

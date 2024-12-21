package iteration

const repeatCount = 5

func Repeat(character string) (repeated string) {
	for range repeatCount {
		repeated += character
	}
	return repeated
}

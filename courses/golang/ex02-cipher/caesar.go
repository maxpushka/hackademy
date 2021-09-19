package cipher

type Caesar struct{}

const caesarShift = 3

func NewCaesar() *Caesar {
	return &Caesar{}
}

func caesar(r rune, shift int) rune {
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func (*Caesar) Encode(source string) string {
	encoded := []rune(lowercaseOnly(source))
	for index, value := range encoded {
		encoded[index] = caesar(value, caesarShift)
	}
	return string(encoded)
}

func (*Caesar) Decode(source string) string {
	decoded := []rune(lowercaseOnly(source))
	for index, value := range decoded {
		decoded[index] = caesar(value, -caesarShift)
	}
	return string(decoded)
}

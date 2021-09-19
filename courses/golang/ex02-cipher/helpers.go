package cipher

func lowercaseOnly(s string) string {
	var b string
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			b += string(v + 32)
		} else if v >= 'a' && v <= 'z' {
			b += string(v)
		}
	}
	return b
}

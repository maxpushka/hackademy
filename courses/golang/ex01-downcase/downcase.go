package downcase

func Downcase(s string) (string, error) {
	var res string

	for _, ch := range s {
		if ch >= 65 && ch <= 90 {
			res += string(ch + 32)
		} else {
			res += string(ch)
		}
	}

	return res, nil
}

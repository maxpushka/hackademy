package cipher

import "strings"

type Vigenere struct {
	key string
}

func NewVigenere(key string) *Vigenere {
	if !testVigenereKey(key) {
		return nil
	}
	return &Vigenere{key}
}

func testVigenereKey(k string) bool {
	if len(k) < 2 ||
		strings.Count(k, string(k[0])) == len(k) ||
		lowercaseOnly(k) != k {
		return false
	}
	return true
}

func letterOffset(r rune) int {
	return int(r - 'a')
}

func (v *Vigenere) Encode(source string) string {
	sourceCleaned := lowercaseOnly(source)
	var encoded string
	keyLength := len(v.key)
	for i, r := range sourceCleaned {
		offset := (letterOffset(r) + letterOffset(rune(v.key[i%keyLength]))) % alphabetSize
		encoded += string(rune('a' + offset))
	}
	return encoded
}

func (v *Vigenere) Decode(source string) string {
	var decoded string
	keyLength := len(v.key)
	for i, r := range source {
		offset := (letterOffset(r) - letterOffset(rune(v.key[i%keyLength]))) % alphabetSize
		if offset < 0 {
			offset += alphabetSize
		}
		decoded += string(rune('a' + offset))
	}
	return decoded
}

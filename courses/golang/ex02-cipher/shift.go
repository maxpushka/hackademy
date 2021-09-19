package cipher

import "strings"

type Shift struct {
	shift int
}

func NewShift(shift int) *Shift {
	if shift == 0 || shift >= alphabetSize || shift <= -alphabetSize {
		return nil
	}
	return &Shift{shift}
}

func (s *Shift) Encode(source string) string {
	sourceCleaned := lowercaseOnly(source)
	result := strings.Map(func(r rune) rune {
		return caesar(r, s.shift)
	}, sourceCleaned)
	return result
}

func (s *Shift) Decode(source string) string {
	sourceCleaned := lowercaseOnly(source)
	result := strings.Map(func(r rune) rune {
		return caesar(r, -s.shift)
	}, sourceCleaned)
	return result
}

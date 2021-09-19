package brackets

func matches(f, s string) bool {
	switch f + s {
	case "{}", "[]", "()":
		return true
	default:
		return false
	}
}

func Bracket(brackets string) (bool, error) {
	var bracketsStack StringStack
	for _, v := range brackets {
		switch v {
		case '{', '[', '(':
			bracketsStack.Push(string(v))
		case '}', ']', ')':
			s := bracketsStack.Pop()
			if !matches(s, string(v)) {
				return false, nil
			}
		}
	}
	return bracketsStack.Size() == 0, nil
}

package stack

type Stack struct {
	data []interface{}
}

func New() *Stack {
	return &Stack{
		data: make([]interface{}, 1),
	}
}

func (s *Stack) Push(elem interface{}) {
	s.data = append(s.data, elem)
}

func (s *Stack) Pop() interface{} {
	if s.Size() == 0 {
		return ""
	}

	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res
}

func (s *Stack) Size() int {
	return len(s.data)
}

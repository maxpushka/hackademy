package brackets

import stack "github.com/maxpushka/hackademy/courses/golang/ex03-stack"

type StringStack struct {
	stack.Stack
}

func (s *StringStack) Push(n string) { s.Stack.Push(n) }
func (s *StringStack) Pop() string   { return s.Stack.Pop().(string) }
func (s *StringStack) Size() int     { return s.Stack.Size() }

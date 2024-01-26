package stack

type Stack struct {
	stack []string
}

// Push 入栈
func (s *Stack) Push(element string) {
	s.stack = append(s.stack, element)
}

// Len 栈的长度
func (s *Stack) Len() int {
	return len(s.stack)
}

// IsEmpty 栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

// Pop 出栈
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	res := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return res
}

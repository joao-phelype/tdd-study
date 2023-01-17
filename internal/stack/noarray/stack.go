package noarray

type Stack struct {
	size     int
	value    string
	previous *Stack
}

func NewStack() *Stack {
	return &Stack{0, "", nil}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Add(value string) {
	previous := Stack{s.size, s.value, s.previous}
	*s = Stack{s.size + 1, value, &previous}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Pop() string {
	value := s.value
	*s = *s.previous
	return value
}

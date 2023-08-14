package stack

type Stack struct {
	length int
	data   []int
}

func New(len int) *Stack {
	return &Stack{
		length: 0,
		data:   make([]int, len),
	}
}

func (s *Stack) Push(value int) {
	s.data[s.length] = value
	s.length += 1
}

func (s *Stack) Pop() int {
	value := s.data[s.length-1]
	s.length -= 1
	return value
}

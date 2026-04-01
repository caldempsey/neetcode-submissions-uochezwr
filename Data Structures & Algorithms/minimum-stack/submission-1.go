type MinStack struct {
	main []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.main = append(s.main, val)
	if len(s.mins) == 0 || val <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	}
}

func (s *MinStack) Pop() {
	top := s.main[len(s.main)-1]
	s.main = s.main[:len(s.main)-1]
	if top == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
}

func (s *MinStack) Top() int {
	return s.main[len(s.main)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}
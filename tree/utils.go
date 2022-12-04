package tree

type Stack struct {
	value []any
}

func (s *Stack) Push(val any) {
	s.value = append(s.value, val)
}

func (s *Stack) Pop() (val any) {
	length := s.Length()
	val = s.value[length-1]
	s.value = s.value[0 : length-1]
	return
}

func (s *Stack) Length() int {
	return len(s.value)
}

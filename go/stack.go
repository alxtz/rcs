package ds

type Stack[T any] struct {
	Slice []T
}

func (s *Stack[T]) Pop() T {
	var itemToPop = s.Slice[len(s.Slice)-1]
	s.Slice = s.Slice[:len(s.Slice)-1]

	return itemToPop
}

func (s *Stack[T]) Push(val T) {
	s.Slice = append(s.Slice, val)
}

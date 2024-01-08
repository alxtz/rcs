package mapset

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Set[T comparable] interface {
	Add(val T) bool
	Cardinality() int
	IsEmpty() bool
	Remove(i T)
	Pop() (T, bool)
	String() string
	MarshalJSON() ([]byte, error)
}

type threadUnsafeSet[T comparable] map[T]struct{}

func NewThreadUnsafeSet[T comparable](vals ...T) Set[T] {
	s := newThreadUnsafeSetWithSize[T](len(vals))
	for _, item := range vals {
		s.Add(item)
	}
	return s
}

func newThreadUnsafeSetWithSize[T comparable](cardinality int) threadUnsafeSet[T] {
	return make(threadUnsafeSet[T], cardinality)
}

func (s threadUnsafeSet[T]) Add(v T) bool {
	prevLen := len(s)
	s[v] = struct{}{}
	return prevLen != len(s)
}

func (s threadUnsafeSet[T]) Cardinality() int {
	return len(s)
}

func (s threadUnsafeSet[T]) IsEmpty() bool {
	return s.Cardinality() == 0
}
func (s threadUnsafeSet[T]) String() string {
	items := make([]string, 0, len(s))

	for elem := range s {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}
func (s threadUnsafeSet[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Cardinality())

	for elem := range s {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

func (s threadUnsafeSet[T]) Pop() (v T, ok bool) {
	for item := range s {
		delete(s, item)
		return item, true
	}
	return v, false
}

func (s threadUnsafeSet[T]) Remove(v T) {
	delete(s, v)
}

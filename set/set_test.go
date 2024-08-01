package set

import "testing"

func TestSetx_Add(t *testing.T) {
	Addvals := []int{1, 2, 3, 1}
	s := NewMapSet[int](10)
	t.Run("Add", func(t *testing.T) {
		for _, val := range Addvals {
			s.Add(val)
		}

	})
}

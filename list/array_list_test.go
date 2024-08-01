package list

import "testing"

func TestArrayList_PopBack(t *testing.T) {
	arr := NewArrayList[int](0)
	arr.PushBack(1)
	arr.PushBack(2)
	arr.PushBack(3)
	arr.PushFront(4)
	arr.PopBack()
	err := arr.Range(func(key, val int) error {
		println(key, val)
		return nil
	})
	if err != nil {
		return
	}
}

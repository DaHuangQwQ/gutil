package list

import (
	"github.com/DaHuangQwQ/gutil/internal/errs"
	"github.com/DaHuangQwQ/gutil/internal/slice"
)

type ArrayList[T any] struct {
	val []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		val: make([]T, 0, cap),
	}
}

func (a *ArrayList[T]) Len() int {
	return len(a.val)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.val)
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	l := a.Len()
	if index < 0 || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	return a.val[index], e
}

func (a *ArrayList[T]) Add(index int, t T) (err error) {
	a.val, err = slice.Add(a.val, t, index)
	return
}

func (a *ArrayList[T]) PushBack(val T) {
	a.val = append(a.val, val)
}

func (a *ArrayList[T]) PushFront(val T) {
	err := a.Add(0, val)
	if err != nil {
		return
	}
}

// Delete 方法会在必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(a.val, index)
	if err != nil {
		return t, err
	}
	a.val = res
	a.shrink()
	return t, nil
}

func (a *ArrayList[T]) Front() T {
	return a.val[0]
}

func (a *ArrayList[T]) Back() T {
	return a.val[len(a.val)-1]
}

func (a *ArrayList[T]) PopBack() {
	_, err := a.Delete(len(a.val) - 1)
	if err != nil {
		return
	}
}

func (a *ArrayList[T]) PopFront() {
	_, err := a.Delete(0)
	if err != nil {
		return
	}
}

// shrink 数组缩容
func (a *ArrayList[T]) shrink() {
	a.val = slice.Shrink(a.val)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range a.val {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.val))
	copy(res, a.val)
	return res
}

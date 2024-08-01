package slice

import "github.com/DaHuangQwQ/gutil/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	res := src[index]
	//从index位置开始，后面的元素依次往前挪1个位置
	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	//去掉最后一个重复元素
	src = src[:length-1]
	return src, res, nil
}

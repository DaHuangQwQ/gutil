package converter

type Converter[Src any, Dst any] interface {
	Convert(src Src) (Dst, error)
}

type ConverterFunc[Src any, Dst any] func(src Src) (Dst, error)

func (cf ConverterFunc[Src, Dst]) Convert(src Src) (Dst, error) {
	return cf(src)
}

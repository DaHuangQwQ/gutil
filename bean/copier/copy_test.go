package copier

import (
	"fmt"
	"testing"
)

type Source struct {
	Field1 string
	Field2 int
	Field3 float64
}

type Destination struct {
	Field1 string
	Field2 int
	Field3 float64
}

func TestConvertField(t *testing.T) {
	var copier Copier[Source, Destination]

	src := &Source{
		Field1: "test",
		Field2: 123,
		Field3: 45.67,
	}

	// 忽略 Field2
	opts := IgnoreFields("Field2")

	dst, err := copier.Copy(src, opts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Copied object: %+v\n", dst)
}

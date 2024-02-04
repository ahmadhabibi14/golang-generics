package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Ahmad",
		Second: "Habibi",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Ahmad",
		Second: "Habibi",
	}

	fmt.Println("Before change:", data.First)

	assert.Equal(t, "Rizky", data.ChangeFirst("Rizky"))
	assert.Equal(t, "Hello Habibi", data.SayHello("Habibi"))

	fmt.Println("After change:", data.First)
}

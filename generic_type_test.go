package golanggenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Ahmad", "Rizky", "Nusantara", "Habibi"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{111, 222, 333, 444, 555}
	PrintBag(numbers)
}

package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}](first, second T) T {
	fmt.Printf("1st param \t: %v | Data type: %T\n", first, first)
	fmt.Printf("2nd param \t: %v | Data type: %T\n", second, second)
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, int64(100), FindMin[int64](int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	fmt.Printf("First value\t: %v\n", first)
	fmt.Printf("Data type\t: %T\n", first)
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Ahmad", "Rizky", "Nusantara", "Habibi",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Ahmad", first)
}

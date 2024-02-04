package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	fmt.Printf("1st param \t: %v | Data type: %T\n", first, first)
	fmt.Printf("2nd param \t: %v | Data type: %T\n", second, second)
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100.0), Min[float64](float64(100.0), float64(200.0)))
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))
	assert.Equal(t, float64(100.0), Min(100.0, 200.0))
	assert.Equal(t, Age(19), Min(Age(19), Age(20)))
}

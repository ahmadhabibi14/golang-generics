package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Printf("Parameter: %v, Data Type: %T\n", param, param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("habi")
	fmt.Println(result)

	var resultNumber int = Length[int](100)
	fmt.Println(resultNumber)
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Habibi")
	assert.Equal(t, "Habibi", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}

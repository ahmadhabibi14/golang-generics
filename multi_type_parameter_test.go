package golanggenerics

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Printf("Parameter: %v, Data Type: %T\n", param1, param1)
	fmt.Printf("Parameter: %v, Data Type: %T\n", param2, param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Habibi", 19)
}

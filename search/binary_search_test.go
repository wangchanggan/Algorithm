package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(data, 2))
}

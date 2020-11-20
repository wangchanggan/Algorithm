package sort

import (
	"fmt"
	"testing"
)

type IntSlice []int

func (p IntSlice) Len() int                         { return len(p) }
func (p IntSlice) Less(i, j int) bool               { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)                    { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Get(i int) (data interface{})     { return p[i] }
func (p IntSlice) Coverage(i int, data interface{}) { p[i] = data.(int) }
func (p IntSlice) LessData(iData, jData interface{}) bool { return iData.(int) < jData.(int) }

func TestBubbleSort(t *testing.T) {
	nums := IntSlice{3,5,1,6,2,9}
	BubbleSort(nums, 0, 5)
	fmt.Println(nums)
}

func TestQuickSort(t *testing.T) {
	nums := IntSlice{3,5,1,6,2,9}
	QuickSort(nums, 0, 5)
	fmt.Println(nums)
}

func TestInsertionSort(t *testing.T) {
	nums := IntSlice{3,5,1,6,2,9}
	InsertionSort(nums, 0, 5)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	nums := IntSlice{3,5,1,6,2,9}
	HeapSort(nums, 0, 5)
	fmt.Println(nums)
}

func TestMergeSort(t *testing.T) {
	nums := IntSlice{3,5,1,6,2,9}
	tmp := make(IntSlice, nums.Len())
	MergeSort(nums, tmp, 0, 5)
	fmt.Println(nums)
}

func TestShellSort(t *testing.T) {
	nums := IntSlice{3,5,1,6,2,9}
	ShellSort(nums)
	fmt.Println(nums)
}
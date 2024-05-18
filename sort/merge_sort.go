package sort

/**
 * 归并排序
 * 时间复杂度：
 *  * 最好：O(nlog2n)
 *  * 最坏：O(nlog2n)
 *  * 平均：O(nlog2n)
 * 空间复杂度：O(n)
 * 稳定性：稳定
 */
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	// 切分数组为两半，并递归对切分后的数组进行排序
	leftNums := MergeSort(nums[:mid])
	rightNums := MergeSort(nums[mid:])
	return merge(leftNums, rightNums)
}

func merge(leftNums, rightNums []int) []int {
	merged := make([]int, 0)
	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(leftNums) && rightIndex < len(rightNums) {
		if leftNums[leftIndex] <= rightNums[rightIndex] {
			merged = append(merged, leftNums[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, rightNums[rightIndex])
			rightIndex++
		}
	}
	merged = append(merged, leftNums[leftIndex:]...)
	merged = append(merged, rightNums[rightIndex:]...)
	return merged
}

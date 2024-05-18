package sort

/**
 * 快速排序
 * 时间复杂度：
 *  * 最好：O(n log2 n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n log2 n)
 * 空间复杂度：O(n log2 n)
 * 稳定性：不稳定
 */
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	// 以第一个为基准值
	pivotIndex := start
	pivotValue := nums[start]
	left, right := start, end
	for left < right {
		// 从小到大排序要先找出比基准值大的
		for left < right && nums[right] > pivotValue {
			right--
		}
		// 找出比基准值小的
		for left < right && nums[left] <= pivotValue {
			left++
		}
		// 交换
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	// 调整基准值，基准值左边的元素均小于等于基准值，右边的元素均大于基准值
	nums[pivotIndex], nums[left] = nums[left], nums[pivotIndex]
	quickSort(nums, start, left-1)
	quickSort(nums, left+1, end)
}

package sort

/**
 * 选择排序
 * 时间复杂度：
 *  * 最好：O(n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n^2)
 * 空间复杂度：O(1)
 * 稳定性：稳定
 */
func SelectionSort(nums []int) {
	// 遍历数组，只需要遍历到倒数第二个元素，因为当只剩下一个元素时，数组已经是有序的了
	for i := 0; i < len(nums)-1; i++ {
		// 将当前位置i作为最小元素的起始位置
		minIndex := i
		// 从i+1开始向后查找比当前minIndex位置的元素更小的元素的位置
		for j := i + 1; j < len(nums); j++ {
			// 如果找到一个更小的元素，更新minIndex
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[minIndex], nums[i] = nums[i], nums[minIndex]
		}
	}
}

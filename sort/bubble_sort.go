package sort

/**
 * 冒泡排序
 * 时间复杂度：
 *  * 最好：O(n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n^2)
 * 空间复杂度：O(1)
 * 稳定性：稳定
 */
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		// 从数组的开始遍历到倒数第二个元素，因为最后一个元素上一轮遍历已经确定是最大，不需要比较
		for j := 0; j < len(nums)-i-1; j++ {
			// 将每一对相邻元素进行比较，较大的元素到后面
			if nums[j] > nums[j+1] {
				//如果当前元素比下一个元素大。则交换它们的位置
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

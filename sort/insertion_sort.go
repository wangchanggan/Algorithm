package sort

/**
 * 插入排序
 * 时间复杂度：
 *  * 最好：O(n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n^2)
 * 空间复杂度：O(1)
 * 稳定性：稳定
 */
func InsertionSort(nums []int) {
	// 从第二个元素开始，默认第一个元素是已经排好序的
	for i := 1; i < len(nums); i++ {
		// 记录当前值，并从当前位置开始向前比较
		key := nums[i]
		// 设置一个变量，跟踪前一个元素的索引
		j := i - 1
		// 当前值key如果小于它前面的值，则进行交换，此过程持续找到合适的插入点
		for j >= 0 && key < nums[j] {
			// 把大的值往后移动
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

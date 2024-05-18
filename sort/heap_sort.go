package sort

/**
 * 堆排序
 * 时间复杂度：
 *  * 最好：O(nlog2 n)
 *  * 最坏：O(nlog2 n)
 *  * 平均：O(nlog2 n)
 * 空间复杂度：O(1)
 * 稳定性：不稳定
 */
func HeapSort(nums []int) {
	heapSize := len(nums)
	// 构建大顶堆
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(nums, len(nums), i)
	}
	// 弹出堆顶元素，交换至序列最后，堆长度-1，然后重新堆化
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, i, 0)
	}
}

func maxHeapify(nums []int, heapLength, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2
	// 判断左子节点是否存在，且大于根节点
	if left < heapLength && nums[left] > nums[largest] {
		largest = left
	}
	// 判断右子节点是否存在，且大于根节点
	if right < heapLength && nums[right] > nums[largest] {
		largest = right
	}
	// 如果跟节点i不是最大的，与最大的替换
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		// 递归堆化
		maxHeapify(nums, heapLength, largest)
	}
}

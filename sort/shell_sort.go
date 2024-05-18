package sort

/**
 * 希尔排序
 * 时间复杂度：
 *  * 最好：O(n^1.3)
 *  * 最坏：O(n^2)
 *  * 平均：O(nlog2 n)
 * 空间复杂度：O(1)
 * 稳定性：不稳定
 */
func ShellSort(nums []int) {
	gap := len(nums) / 2
	for gap > 0 {
		for i := gap; i < len(nums); i++ {
			tmp := nums[i]
			j := i
			for j >= gap && nums[j-gap] > tmp {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = tmp
		}
		gap /= 2
	}
}

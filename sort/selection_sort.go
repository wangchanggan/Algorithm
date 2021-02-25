package sort

/**
 * 选择排序
 * 时间复杂度：
 *  * 最好：O(n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n^2)
 * 空间复杂度：O(1)
 * 稳定性：稳定
 * @param SortInterface data 源数据
 * @param int start 需排序起始位置
 * @param int end 需排序末尾位置
 */
func SelectionSort(data SortInterface, start, end int) {
	if data.Len() <= 1 {
		return
	}

	for i := start; i < end; i++ {
		index := i
		for j := i + 1; j <= end; j++ {
			if data.Less(index, j) {
				index = j
			}
		}
		data.Swap(index, i)
	}
}

package sort

/**
 * 插入排序
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
func InsertionSort(data SortInterface, start, end int) {
	if data.Len() <= 1 {
		return
	}

	for i := start + 1; i <= end; i++ {
		for j := i; j > start && data.Less(j-1, j); j-- {
			data.Swap(j, j-1)
		}
	}
}
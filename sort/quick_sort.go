package sort

/**
 * 快速排序
 * 时间复杂度：
 *  * 最好：O(n log2 n)
 *  * 最坏：O(n^2)
 *  * 平均：O(n log2 n)
 * 空间复杂度：O(n log2 n)
 * 稳定性：不稳定
 * @param SortInterface data 源数据
 * @param int start 需排序起始位置
 * @param int end 需排序末尾位置
 */
func QuickSort(data SortInterface, start, end int) {
	if data.Len() <= 1 {
		return
	}

	if start < end && data.Len() > start {
		//以第一个为基准值
		pivotValue := data.Get(start).(int)
		s := start
		e := end

		for s < e {
			for s < e && data.LessData(data.Get(e).(int), pivotValue) {
				e--
			}
			if s < e {
				data.Coverage(s, data.Get(e).(int))
				s++
			}

			for s < e && !data.LessData(data.Get(s).(int), pivotValue) {
				s++
			}
			if s < e {
				data.Coverage(e, data.Get(s).(int))
				e--
			}
		}
		data.Coverage(s, pivotValue)
		QuickSort(data, start, s-1)
		QuickSort(data, s+1, end)
	}
}
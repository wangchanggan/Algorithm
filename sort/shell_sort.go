package sort

/**
 * 希尔排序
 * 时间复杂度：
 *  * 最好：O(n^1.3)
 *  * 最坏：O(n^2)
 *  * 平均：O(nlog2 n)
 * 空间复杂度：O(1)
 * 稳定性：不稳定
 * @param SortInterface data 源数据
 * @param int start 需排序起始位置
 * @param int end 需排序末尾位置
 */
func ShellSort(data SortInterface) {
	if data.Len() <= 1 {
		return
	}

	for step := data.Len() / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < data.Len(); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && data.Less(j, j+step); j -= step {
				data.Swap(j, j+step)
			}
		}
	}
}

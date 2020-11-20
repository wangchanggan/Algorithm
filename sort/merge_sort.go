package sort

/**
 * 归并排序
 * 时间复杂度：
 *  * 最好：O(nlog2n)
 *  * 最坏：O(nlog2n)
 *  * 平均：O(nlog2n)
 * 空间复杂度：O(n)
 * 稳定性：稳定
 * @param SortInterface data 源数据
 * @param SortInterface tmp 临时暂存数据
 * @param int start 需排序起始位置
 * @param int end 需排序末尾位置
 */
func MergeSort(data, tmp SortInterface, start, end int) {
	if data.Len() <= 1 {
		return
	}

	if start < end {
		mid := start + (end-start)/2
		MergeSort(data, tmp, start, mid);
		MergeSort(data, tmp, mid+1, end);
		merge(data, tmp, start, mid, end);
	}
}

func merge(data, tmp SortInterface, start, mid, end int) {
	i := start
	j := mid + 1
	k := start

	for i != mid+1 && j != end+1 {
		if data.Less(i, j) {
			tmp.Coverage(k, data.Get(j))
			k++
			j++
		} else {
			tmp.Coverage(k, data.Get(i))
			k++
			i++
		}
	}

	for i != mid+1 {
		tmp.Coverage(k, data.Get(i))
		k++
		i++
	}

	for j != end+1 {
		tmp.Coverage(k, data.Get(j))
		k++
		j++
	}

	for i := start; i <= end; i++ {
		data.Coverage(i, tmp.Get(i))
	}
}

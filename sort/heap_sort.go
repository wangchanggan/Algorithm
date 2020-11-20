package sort

/**
 * 堆排序
 * 时间复杂度：
 *  * 最好：O(nlog2 n)
 *  * 最坏：O(nlog2 n)
 *  * 平均：O(nlog2 n)
 * 空间复杂度：O(1)
 * 稳定性：不稳定
 * @param SortInterface data 源数据
 * @param int start 需排序起始位置
 * @param int end 需排序末尾位置
 */
func HeapSort(data SortInterface, start, end int) {
	if data.Len() <= 1 {
		return
	}

	first := start
	lo := 0
	hi := end + 1 - start

	// 构建大顶堆.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		// 将堆顶元素与末尾元素交换，将最大元素"沉"到数组末端
		data.Swap(first, first+i)
		// 重新调整结构，使其满足大顶堆
		siftDown(data, lo, i, first)
	}
}

func siftDown(data SortInterface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && !data.Less(first+child, first+child+1) {
			child++
		}
		if data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}
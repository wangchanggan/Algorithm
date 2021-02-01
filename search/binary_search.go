package search

func binarySearchInternally(data []int, low, high int, value int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)>>1 // low + (high - low)/2
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return binarySearchInternally(data, mid+1, high, value)
	} else {
		return binarySearchInternally(data, low, mid-1, value)
	}
}

/**
 * 二分查找
 * @param []int data 有序的源数据
 * @param int value 查找的数据
 * @return int index 查找数据的下标
 */
func BinarySearch(data []int, value int) (index int) {
	return binarySearchInternally(data, 0, len(data)-1, value)
}

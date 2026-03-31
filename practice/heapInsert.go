package practice

// heapInsert 向上调整
// arr: 原始数组
// index: 当前需要向上调整的元素的下标
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

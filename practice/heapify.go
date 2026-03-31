package practice

// heapify 向下调整
// arr: 原始数组
// index: 当前需要向下调整的元素的下标
// heapSize: 当前堆的大小（用来控制边界，判断有没有越界）
func heapify(arr []int, index int, heapSize int) {
	// 左孩子的下标是 2*index + 1
	left := index*2 + 1

	for left < heapSize {
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}

		if arr[index] >= arr[largest] {
			break
		}

		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
		left = index*2 + 1
	}
}

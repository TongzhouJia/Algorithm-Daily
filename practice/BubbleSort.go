package practice

// BubbleSort 使用冒泡排序算法对整数切片进行原地升序排列。
func BubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				// Go 特色：不需要 temp 变量，直接交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

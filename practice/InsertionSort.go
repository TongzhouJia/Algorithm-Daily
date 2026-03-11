package practice

// InsertionSort 使用插入排序算法对整数切片进行原地升序排列。
// 1. 为什么循环从 i := 1 开始？
// 因为单个元素（索引 0）默认是有序的，从第二个元素开始扫描并插入到前面已排序序列的合适位置。
// 2. 边界情况判定：
// 如果 n < 2判定为假，循环不会执行，自然处理了空切片或单元素切片。
// 3. arr[j+1] = key用到是j+1是因为在最后一轮循环多减了1
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// 将比key大的元素向后移动一位，为key腾出空间
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// 插入
		arr[j+1] = key
	}
}

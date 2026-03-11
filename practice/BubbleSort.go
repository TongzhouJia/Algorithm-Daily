package practice

// BubbleSort 使用冒泡排序算法对整数切片进行原地升序排列。
// 1. 为什么不做if n<1边界情况判定？
// 因为会在第一个for循环中i<n-1直接判定为假
// 2. arr[j], arr[j+1] = arr[j+1], arr[j]这是go语法糖，不需要temp
// 3. swapped是为了提前结束，如果后边全都是顺序对的了就可以提前跳出
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

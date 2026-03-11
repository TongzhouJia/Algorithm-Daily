package practice

// SelectionSort 对切片进行原地选择排序。
// 选择排序找最小值依次放到前边，所以最前边的一串数组一定是排好序的，每次是在后边数组里选的目前最小的和前边排好序最后一位交换
// 1. 为什么不比较value而是index，如果index是不同的但值是相同的岂不是多交换了？
// 那是因为如果比较index只需要看CPU寄存器里两个数字，但value需要去主存或缓存先取出，而很显然这种极端情况不常见
// 2. 为什么不做if n<1边界情况判定？
// 因为会在第一个for循环中i<n-1直接判定为假
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		// 直接交换以保持代码简洁并减少分支预测开销
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

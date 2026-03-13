package practice

// MergeSort 供外部调用的主方法
func MergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	processMergeSort(arr, 0, len(arr)-1)
}

// processMergeSort 核心递归排序过程
func processMergeSort(arr []int, L int, R int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	// 让左半部分有序
	processMergeSort(arr, L, mid)
	// 让右半部分有序
	processMergeSort(arr, mid+1, R)

	// 左右合并，让整体有序
	merge(arr, L, mid, R)
}

// merge 左右两侧合并的具体实现
func merge(arr []int, L int, mid int, R int) {
	// 开辟等规模的辅助空间
	help := make([]int, R-L+1)
	i := 0 // help数组的专用指针
	p1 := L
	p2 := mid + 1

	// 当左侧和右侧都没越界时，谁小拷贝谁
	for p1 <= mid && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	// 下面两个for循环只会执行其中一个
	// 如果p1没越界把p1剩下的数全拷进去
	for p1 <= mid {
		help[i] = arr[p1]
		i++
		p1++
	}

	// 如果p2没越界，把p2剩下的数全拷进去
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}

	// 将help数组的内容拷贝回原数组对应的[L,R]范围
	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}
}

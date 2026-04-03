package practice

// SmallSum 供外部调用的主方法
func SmallSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	return processSmallSum(arr, 0, len(arr)-1)
}

// processSmallSum 递归求解并排序
func processSmallSum(arr []int, L int, R int) int {
	if L == R {
		return 0
	}

	mid := L + ((R - L) >> 1)

	// 整体的小和=左侧部分产生的小和+右侧部分产生的小和+跨越左右两侧merge时产生的小和
	return processSmallSum(arr, L, mid) +
		processSmallSum(arr, mid+1, R) +
		mergeSmallSum(arr, L, mid, R)
}

// mergeSmallSum 在合并过程中计算小和
func mergeSmallSum(arr []int, L int, mid int, R int) int {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid + 1
	res := 0 // 累加小和的变量

	for p1 <= mid && p2 <= R {
		if arr[p1] < arr[p2] {
			res += (R - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		help[i] = arr[p1]
		i++
		p1++
	}

	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}

	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}

	return res
}

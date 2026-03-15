package practice

// ReversePairs 供外部调用的主方法
func ReversePairs(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return processReversePairs(arr, 0, len(arr)-1)
}

// processReversePairs 递归求解并排序
func processReversePairs(arr []int, L int, R int) int {
	if L == R {
		return 0
	}

	mid := L + ((R - L) >> 1)

	return processReversePairs(arr, L, mid) +
		processReversePairs(arr, mid+1, R) +
		mergeReversePairs(arr, L, mid, R)
}

// mergeReversePairs 在合并过程中计算逆序对
func mergeReversePairs(arr []int, L int, mid int, R int) int {
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := mid + 1
	res := 0 // 累加逆序对的变量

	for p1 <= mid && p2 <= R {
		// 相等时优先拷贝左边，不产生逆序对
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			// 左侧大于右侧，p1到mid的数都跟arr[p2]构成逆序对
			res += mid - p1 + 1
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

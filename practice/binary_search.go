package practice

// BinarySearchExist 在一个有序数组中，找某个数是否存在 (最基础的二分查找)
// 1. left和right对于mid加一减一避免了死循环
func BinarySearchExist(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == num {
			return true // 找到了，直接返回
		} else if arr[mid] < num {
			left = mid + 1 // num在右侧，左边往右移
		} else {
			right = mid - 1 // num在左侧，右边往左移
		}
	}
	return false
}

// BinarySearchNearestLeft 在一个有序数组中，找>=某个数最左侧的位置
func BinarySearchNearestLeft(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	ans := -1 //临时记录，用个肯定不可能的数

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= num {
			ans = mid       // 满足条件，记录下来
			right = mid - 1 // 接着缩小范围
		} else {
			left = mid + 1 // 不满足条件，说明目标还在右边
		}
	}

	return ans
}

// BinarySearchLocalMin 局部最小值问题（无序数组）找到“⬇️最小值⬆️”，时间复杂度小于O（n）
func BinarySearchLocalMin(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}

	// 找趋势，一个下降一个上升，0，1是下降趋势，N-2，N-1上升的
	// 局部最小必然在 [1, N-2] 之间
	left, right := 1, len(arr)-2

	for left <= right {
		mid := left + (right-left)/2

		// 如果 mid 处正好比左右都小，它就是一个“坑”
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		}

		// 先判断左边，如果左边是下降，说明左边必有局部最小
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else {
			// 否则右边是下降的，右边必有局部最小
			left = mid + 1
		}
	}

	return left
}

package practice

// MaxDivideConquer 递归获得最大值
func MaxDivideConquer(nums []int) int {
	if len(nums) == 0 {
		// 根据实际工程要求，空切片求最大值可能需要 panic 或返回特定错误码
		// 这里作为算法练习的兜底，暂时返回 0
		return 0
	}
	// 调用包内私有的递归辅助函数
	return maxRecursive(nums, 0, len(nums)-1)
}

// maxRecursive内部核心递归函数
// 小驼峰命名私有函数
//
//	mid := l + (r - l) >> 1可以不使用外层括号因为位运算优先级更高
func maxRecursive(nums []int, l, r int) int {
	// 递归的出口（Base Case）：范围缩小到只有一个元素时，直接返回
	if l == r {
		return nums[l]
	}
	// 求中点：防止整数溢出，且利用位运算加速
	mid := l + (r-l)>>1
	// 分治：分别求左右两侧的最大值
	leftMax := maxRecursive(nums, l, mid)
	rightMax := maxRecursive(nums, mid+1, r)

	// 决策：合并左右两侧的结果
	return max(leftMax, rightMax)
}

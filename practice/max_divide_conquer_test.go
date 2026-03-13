package practice

import (
	"math/rand"
	"testing"
)

// TestMaxDivideConquer_Basic 表格驱动测试，用于覆盖常见的边界和特殊情况
func TestMaxDivideConquer_Basic(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "nil 切片",
			nums: nil,
			want: 0, // 你的代码中 len(nums) == 0 返回 0
		},
		{
			name: "空切片",
			nums: []int{},
			want: 0,
		},
		{
			name: "单元素切片",
			nums: []int{42},
			want: 42,
		},
		{
			name: "正数切片",
			nums: []int{1, 5, 3, 9, 2},
			want: 9,
		},
		{
			name: "负数切片",
			nums: []int{-5, -2, -9, -1, -3},
			want: -1,
		},
		{
			name: "包含正负数和0",
			nums: []int{-5, 2, 0, 8, -3},
			want: 8,
		},
		{
			name: "全部是相同的元素",
			nums: []int{7, 7, 7, 7, 7},
			want: 7,
		},
		{
			name: "最大值在首位",
			nums: []int{100, 2, 3, 4, 5},
			want: 100,
		},
		{
			name: "最大值在末位",
			nums: []int{1, 2, 3, 4, 100},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxDivideConquer(tt.nums)
			if got != tt.want {
				t.Errorf("MaxDivideConquer() = %v, 期望得到 %v", got, tt.want)
			}
		})
	}
}

// TestMaxDivideConquer_Random 对数器测试，生成大量随机数组与标准 O(N) 遍历比较
func TestMaxDivideConquer_Random(t *testing.T) {

	testTimes := 5000 // 随机测试的次数
	maxSize := 500    // 随机数组的最大长度
	maxValue := 10000 // 随机数组元素的最大绝对值

	for i := 0; i < testTimes; i++ {
		// 1. 生成随机数组
		arr := generateRandomArrayForMax(maxSize, maxValue)

		// 2. 分别使用你的分治方法和标准的遍历方法求最大值
		got := MaxDivideConquer(arr)
		want := standardMax(arr)

		// 3. 对比结果
		if got != want {
			t.Fatalf("随机测试失败！\n测试数组: %v\n你的结果: %v\n正确结果: %v\n", arr, got, want)
		}
	}
}

// standardMax 对数器的绝对正确标准方法（O(N) 普通遍历求最大值）
func standardMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// generateRandomArrayForMax 辅助方法：生成随机数组（复用或略微调整上一份测试的数据生成器）
func generateRandomArrayForMax(maxSize int, maxValue int) []int {
	size := rand.Intn(maxSize + 1)
	if size == 0 {
		return []int{}
	}

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		// 随机生成 [-maxValue, maxValue] 范围内的数字
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

package practice

import (
	"testing"
)

// TestBinarySearchExist 测试基础二分查找
func TestBinarySearchExist(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		num      int
		expected bool
	}{
		{"存在于中间", []int{1, 3, 5, 7, 9}, 5, true},
		{"存在于开头", []int{1, 3, 5, 7, 9}, 1, true},
		{"存在于结尾", []int{1, 3, 5, 7, 9}, 9, true},
		{"不存在-过小", []int{1, 3, 5, 7, 9}, 0, false},
		{"不存在-过大", []int{1, 3, 5, 7, 9}, 10, false},
		{"不存在-在中间间隔", []int{1, 3, 5, 7, 9}, 4, false},
		{"空数组", []int{}, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchExist(tt.arr, tt.num); got != tt.expected {
				t.Errorf("BinarySearchExist() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestBinarySearchNearestLeft 测试查找 >= num 最左的位置
func TestBinarySearchNearestLeft(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		num      int
		expected int
	}{
		{"正常查找", []int{1, 2, 2, 2, 3, 3, 5}, 2, 1},
		{"查找不存在但有更大的", []int{1, 3, 5, 7}, 4, 2}, // 第一个 >= 4 的是 5，索引为 2
		{"全部满足条件", []int{5, 6, 7}, 2, 0},
		{"全部不满足条件", []int{1, 2, 3}, 5, -1},
		{"空数组", []int{}, 5, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchNearestLeft(tt.arr, tt.num); got != tt.expected {
				t.Errorf("BinarySearchNearestLeft() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestBinarySearchLocalMin 测试局部最小值
func TestBinarySearchLocalMin(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"单元素", []int{1}},
		{"开头是最小", []int{1, 2, 3, 4, 5}},
		{"结尾是最小", []int{5, 4, 3, 2, 1}},
		{"中间凹陷", []int{3, 2, 3, 2, 3}},
		{"复杂波形", []int{9, 7, 10, 8, 11, 6, 12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := BinarySearchLocalMin(tt.arr)

			// 验证返回的索引是否合法
			if len(tt.arr) == 0 {
				if idx != -1 {
					t.Errorf("空数组应返回 -1, 得到 %d", idx)
				}
				return
			}

			// 验证局部最小值的定义
			val := tt.arr[idx]
			leftOK, rightOK := true, true

			if idx > 0 {
				leftOK = val < tt.arr[idx-1]
			}
			if idx < len(tt.arr)-1 {
				rightOK = val < tt.arr[idx+1]
			}

			if !leftOK || !rightOK {
				t.Errorf("索引 %d 处的值 %d 不是局部最小值 (数组: %v)", idx, val, tt.arr)
			}
		})
	}
}

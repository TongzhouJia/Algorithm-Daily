package practice

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "常规乱序数组",
			input:    []int{64, 25, 12, 22, 11},
			expected: []int{11, 12, 22, 25, 64},
		},
		{
			name:     "已经是升序",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "完全逆序",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "空切片",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单个元素",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "包含重复元素",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 拷贝一份数据进行测试，避免影响原始用例（养成好习惯）
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			SelectionSort(input)

			// 检查结果是否符合预期
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("测试项 [%s] 失败: 得到 %v, 想要 %v", tt.name, input, tt.expected)
			}
		})
	}
}

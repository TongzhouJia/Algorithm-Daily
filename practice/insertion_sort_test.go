package practice

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "正常用例",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "空切片",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单个元素",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "已排序切片",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "逆序切片",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "包含重复元素",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 拷贝输入数据，避免原地排序影响原始测试用例数据
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			InsertionSort(input)

			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf("%s 失败: 得到 %v, 想要 %v", tt.name, input, tt.expected)
			}
		})
	}
}

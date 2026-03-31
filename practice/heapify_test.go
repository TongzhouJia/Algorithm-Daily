package practice

import (
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		index    int
		heapSize int
		expected []int
	}{
		{
			name:     "已经是大根堆-无需下沉",
			initial:  []int{10, 8, 9},
			index:    0,
			heapSize: 3,
			expected: []int{10, 8, 9},
		},
		{
			name:     "下沉一层-与左孩子交换",
			initial:  []int{5, 10, 8},
			index:    0,
			heapSize: 3,
			expected: []int{10, 5, 8},
		},
		{
			name:     "下沉一层-与右孩子交换",
			initial:  []int{5, 8, 10},
			index:    0,
			heapSize: 3,
			expected: []int{10, 8, 5},
		},
		{
			name:     "连续下沉多层",
			initial:  []int{1, 10, 9, 5, 4, 3, 2},
			index:    0,
			heapSize: 7,
			expected: []int{10, 5, 9, 1, 4, 3, 2},
		},
		{
			name:     "受限于 heapSize-不处理界外元素",
			initial:  []int{5, 10, 15},
			index:    0,
			heapSize: 2, // 虽然有15，但 heapSize 限制了只能看到 index 0和1
			expected: []int{10, 5, 15},
		},
		{
			name:     "只有左孩子的情况",
			initial:  []int{5, 10},
			index:    0,
			heapSize: 2,
			expected: []int{10, 5},
		},
		{
			name:     "已经是叶子节点-无处可沉",
			initial:  []int{10, 8, 9},
			index:    1,
			heapSize: 3,
			expected: []int{10, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制切片，防止污染测试用例
			arr := make([]int, len(tt.initial))
			copy(arr, tt.initial)

			heapify(arr, tt.index, tt.heapSize)

			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("heapify() 失败 '%s': 实际 %v, 想要 %v", tt.name, arr, tt.expected)
			}
		})
	}
}

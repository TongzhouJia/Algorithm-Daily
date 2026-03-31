package practice

import (
	"reflect"
	"testing"
)

func TestHeapInsert(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		index    int
		expected []int
	}{
		{
			name:     "插入到空堆或作为根节点",
			initial:  []int{10},
			index:    0,
			expected: []int{10},
		},
		{
			name:     "新元素大于父节点-触发交换",
			initial:  []int{10, 5, 8, 15}, // 15 是新插入在 index 3 的元素
			index:    3,
			expected: []int{15, 10, 8, 5},
		},
		{
			name:     "新元素小于父节点-不触发交换",
			initial:  []int{20, 10, 15, 5}, // 5 是新插入在 index 3 的元素
			index:    3,
			expected: []int{20, 10, 15, 5},
		},
		{
			name:     "连续向上交换多次",
			initial:  []int{50, 40, 30, 20, 10, 60}, // 60 插入在 index 5
			index:    5,
			expected: []int{60, 40, 50, 20, 10, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制一份切片，避免修改原始测试用例数据
			arr := make([]int, len(tt.initial))
			copy(arr, tt.initial)

			heapInsert(arr, tt.index)

			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("heapInsert() 失败 = %v, 想要 %v", arr, tt.expected)
			}
		})
	}
}

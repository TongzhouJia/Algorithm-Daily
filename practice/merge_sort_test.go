package practice

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// TestMergeSort_Basic 表格驱动测试，用于覆盖常见的边界和特殊情况
func TestMergeSort_Basic(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "nil 数组",
			arr:  nil,
			want: nil,
		},
		{
			name: "空数组",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "单元素数组",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "已经有序的数组",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "完全逆序的数组",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "包含重复元素的数组",
			arr:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			want: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name: "包含负数的数组",
			arr:  []int{-5, 2, -1, 0, 8, -3},
			want: []int{-5, -3, -1, 0, 2, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input []int
			if tt.arr != nil {
				// 拷贝一份原数组，防止测试用例的数据被直接修改
				input = make([]int, len(tt.arr))
				copy(input, tt.arr)
			}

			MergeSort(input)

			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("MergeSort() = %v, 期望得到 %v", input, tt.want)
			}
		})
	}
}

// TestMergeSort_Random 对数器测试，生成大量随机数组与 Go 官方的排序进行对比
func TestMergeSort_Random(t *testing.T) {

	testTimes := 1000 // 随机测试的次数
	maxSize := 100    // 随机数组的最大长度
	maxValue := 1000  // 随机数组元素的最大绝对值

	for i := 0; i < testTimes; i++ {
		// 1. 生成随机数组
		arr1 := generateRandomArray(maxSize, maxValue)

		// 2. 拷贝一份用于官方排序
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)

		// 3. 分别使用你的方法和官方方法进行排序
		MergeSort(arr1)
		sort.Ints(arr2)

		// 4. 对比结果
		if !reflect.DeepEqual(arr1, arr2) {
			t.Fatalf("随机测试失败！\n你的结果: %v\n正确结果: %v\n", arr1, arr2)
		}
	}
}

// generateRandomArray 辅助方法：生成长度随机、值也随机的数组
func generateRandomArray(maxSize int, maxValue int) []int {
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

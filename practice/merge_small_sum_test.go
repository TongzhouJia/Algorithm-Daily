package practice

import (
	"math/rand"
	"testing"
)

func TestSmallSum_Random(t *testing.T) {
	// 直接在这里定义局部变量函数，不会跟其他文件冲突
	generateRandomArray := func(maxSize, maxValue int) []int {
		size := rand.Intn(maxSize + 1)
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
		}
		return arr
	}

	comparator := func(arr []int) int {
		if len(arr) < 2 {
			return 0
		}
		res := 0
		for i := 1; i < len(arr); i++ {
			for j := 0; j < i; j++ {
				if arr[j] < arr[i] {
					res += arr[j]
				}
			}
		}
		return res
	}

	for i := 0; i < 5000; i++ {
		origin := generateRandomArray(100, 100)
		testArr := make([]int, len(origin))
		copy(testArr, origin)

		got := SmallSum(testArr)
		want := comparator(origin)

		if got != want {
			t.Fatalf("失败！数组:%v, 得到:%v, 期望:%v", origin, got, want)
		}
	}
}

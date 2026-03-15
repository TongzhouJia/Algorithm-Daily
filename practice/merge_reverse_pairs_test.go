package practice

import (
	"math/rand"
	"testing"
)

func TestReversePairs_Random(t *testing.T) {
	// 同样的方法，写死在内部
	gen := func(maxSize, maxValue int) []int {
		size := rand.Intn(maxSize + 1)
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(maxValue + 1)
		}
		return arr
	}

	comp := func(arr []int) int {
		res := 0
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					res++
				}
			}
		}
		return res
	}

	for i := 0; i < 5000; i++ {
		origin := gen(100, 100)
		testArr := make([]int, len(origin))
		copy(testArr, origin)

		got := ReversePairs(testArr)
		want := comp(origin)

		if got != want {
			t.Fatalf("失败！数组:%v, 得到:%v, 期望:%v", origin, got, want)
		}
	}
}

package leetcode

import (
	"fmt"
	"log"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		p := q.para1                    // 输入参数 包含数组明细和target
		ans := twoSum(p.nums, p.target) // 计算出来的答案
		if len(q.ans1.one) != len(ans) {  // q.ans1.one 期待答案 ans计算出来的答案
			fmt.Printf("Not Pass【input】:%v       【output】:%v\n", p, ans)
			log.Fatalf("**** Not Pass!!! ****")
			return
		}
		for i, v := range q.ans1.one {
			if v != ans[i] {
				fmt.Printf("Not Pass【input】:%v       【output】:%v\n", p, ans)
				log.Fatalf("**** Not Pass!!! ****")
				return
			}
		}

	}
	fmt.Printf("------------------------Leetcode Problem 1 All Pass.------------------------\n")
}

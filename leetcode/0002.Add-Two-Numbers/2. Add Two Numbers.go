package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 不能直接计算 每个链表中的节点数在范围 [1, 100] 内 int64也会溢出
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	//var a int64 = 1000000000000000000000000000001
//	//fmt.Println("a=",a)
//	sum1,sum2,idx1,idx2:=0,0,0,0
//	curL1,curL2:=l1,l2
//	for ;curL1!=nil;{
//		sum1+=int(math.Pow10(idx1))*curL1.Val
//		idx1++
//		curL1=curL1.Next
//	}
//	for ;curL2!=nil;{
//		sum2+=int(math.Pow10(idx2))*curL2.Val
//		idx2++
//		curL2=curL2.Next
//	}
//	total:=sum1+sum2
//	if l1==nil && l2==nil {
//		return nil
//	}
//	head := &ListNode{Val: 0}
//	result:=strconv.Itoa(total)
//	var curNode *ListNode
//	for i:=len(result)-1;i>=0;i--{
//		fmt.Println(string(result[i]))
//		var node ListNode
//		node.Val,_= strconv.Atoi(string(result[i]))
//		node.Next=nil
//		if i==len(result)-1{
//			head.Next=&node
//			curNode = &node
//		}else{
//			curNode.Next=&node
//			curNode = &node
//		}
//	}
//	return head.Next
//}
//[]int{1}, []int{9, 9, 9, 9, 9}
// addTwoNumbers 数字都是逆序存放 所以可以直接相加(注意进位),短的链表 补充0即可(跟正序存放前面补0一样) 链接遍历结束如果进位不为0 还要添加一个节点
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 声明一个虚拟头节点 方便code
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	// 2个链表没有都遍历完成 或者 都遍历完成 但是最后进位不是0
	for l1 != nil || l2 != nil || carry != 0 {
		// 1号链表遍历完成 2号链表还未遍历完成，1号链表补0
		if l1 == nil {
			n1 = 0
		} else {
			// n1为1号链表 当前数字 保存后 l1链表后移
			n1 = l1.Val
			l1 = l1.Next
		}
		// 2号链表遍历完成 1号链表还未遍历完成，2号链表补0
		if l2 == nil {
			n2 = 0
		} else {
			// n2为2号链表 当前数字 保存后 l2链表后移
			n2 = l2.Val
			l2 = l2.Next
		}
		// 构造2个链表数字相加后的节点
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		// 链表移到尾部
		current = current.Next
		// 进位
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

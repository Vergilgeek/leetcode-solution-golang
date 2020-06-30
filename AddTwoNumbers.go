package main

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func solution(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	var a = l1
	var b = l2
	var result = new(ListNode)
	var node = result
	for nil != a || nil != b {
		var aValue = 0
		if nil != a {
			aValue = a.Val
			a = a.Next
		}
		var bValue = 0
		if nil != b {
			bValue = b.Val
			b = b.Next
		}
		var temp = aValue + bValue + carry
		var tempNode = new(ListNode)
		if temp < 10 {
			carry = 0
			tempNode.Val = temp
		} else {
			carry = 1
			temp = temp % 10
			tempNode.Val = temp
		}
		node.Next = tempNode
		node = node.Next
	}
	if carry > 0 {
		node.Next = &ListNode{Val: 1}
	}
	return result.Next
}

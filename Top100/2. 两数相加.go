package main

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//可以用head := new(ListNode),tail = head
	var head, tail *ListNode
	//进位
	var carry = 0

	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

//func main() {
//
//	var l1 = new(ListNode)
//	l1.Val = 1
//	var l1n1 = new(ListNode)
//	l1n1.Val = 2
//	l1.Next = l1n1
//	var l1n2 = new(ListNode)
//	l1n2.Val = 2
//	l1n1.Next = l1n2
//
//	var l2 = new(ListNode)
//	l2.Val = 1
//	var l2n1 = new(ListNode)
//	l2n1.Val = 2
//	l2.Next = l1n1
//	var l2n2 = new(ListNode)
//	l2n2.Val = 2
//	l2n1.Next = l2n2
//
//	l3 := addTwoNumbers(l1, l2)
//	for l3 != nil {
//		fmt.Println(*l3)
//		l3 = l3.Next //移动指针
//	}
//
//}

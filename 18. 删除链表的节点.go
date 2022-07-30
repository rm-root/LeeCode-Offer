package main

//给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//返回删除后的链表的头节点。

//双指针法
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	pre, cur := head, head.Next
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}

	if cur != nil {
		pre.Next = cur.Next
	}
	return head
}

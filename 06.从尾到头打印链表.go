package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	cur := head //备份链表
	n := 0
	//遍历链表记录链表数
	for head != nil {
		head = head.Next
		n++
	}
	//再次遍历链表逆序摆放进数组
	nums := make([]int, n) //申请一个数组,n可作用在这里预申请空间,大大减少了内存消耗
	for {
		if cur == nil {
			break
		}
		nums[n-1] = cur.Val
		cur = cur.Next
		n--
	}
	return nums
}

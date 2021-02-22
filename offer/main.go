package main

import "fmt"

func main() {
	tailListNode := &ListNode{Val: 10}

	secondListNode := &ListNode{Val: 9, Next: tailListNode}

	firstListNode := &ListNode{Val: 8, Next: secondListNode}

	ret := findKthToTail(firstListNode, 1)
	fmt.Println(ret.Val)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func findKthToTail(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	cur := head
	for cur != nil && k > 0 {
		cur = cur.Next
		k--
	}
	//如果最后k是大于cur的长度的，就return
	if cur == nil && k > 0 {
		return nil
	}

	ret := head
	for cur != nil {
		cur = cur.Next
		ret = ret.Next
	}

	return ret

}

func findKthToTail2(head *ListNode, k int) *ListNode {
	if head == nil || k < 0 {
		return nil
	}

	cur := head
	n := 0
	for cur != nil {
		n++
		cur = cur.Next
	}

	if n < k {
		return nil
	}

	ret := head
	for n-k > 0 {
		ret = ret.Next
		k++
	}
	return ret
}

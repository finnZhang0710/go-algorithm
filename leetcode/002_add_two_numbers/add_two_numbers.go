package _02_add_two_numbers


/**

	sum := carry % 10 		//取商，  12 / 10 = 1
	carry := carry / 10  //这个是除号	12 % 10 = 2

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, cur *ListNode
	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		sum := carry % 10
		carry = carry / 10

		node := &ListNode{Val: sum}

		// res = cur = node, 第一个初始化的值，后续就放到next里面
		if res == nil {
			res = node
			cur = node
		} else {
			cur.Next = node
			cur = cur.Next
		}

	}

	return res

}

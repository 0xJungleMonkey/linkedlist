/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	var sentinel *ListNode = &ListNode{Next: head}

	var temp = sentinel

	var temp2 = sentinel

	for temp != nil && temp.Next != nil {
		temp = temp.Next.Next
		temp2 = temp2.Next
		fmt.Println("2help!")
	}
	fmt.Println("4help!")
	temp2 = temp2.Next
	fmt.Println("help!")
	var revHead *ListNode
	for temp2 != nil {
		temp2.Next, revHead, temp2 = revHead, temp2, temp2.Next
	}
	for revHead != nil {
		if head.Val == revHead.Val {
			head = head.Next
			revHead = revHead.Next
		} else {
			return false
		}
	}
	return true

}
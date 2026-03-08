package leetcode

type ListNode struct {
     Val int
     Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var new_head *ListNode;
	var pre_head *ListNode;
	for head.Next != nil {
		new_head = head.Next;
		head.Next, pre_head = pre_head, head
		head = new_head
	}
	if pre_head != nil {
		head.Next = pre_head
	}
	return head
    
}
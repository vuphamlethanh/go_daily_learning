package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func CalNumListNode(l *ListNode) int {
	result := 0
	length := 0

	for l.Next != nil {
		length += 1
		result += l.Val ^ length
		l = l.Next
	}

	return result
}

func NumToListNode(num int) *ListNode {
	count := 0

	var listNode *ListNode = nil

	for {
		count++

		val := (num%10^count)/1 ^ count
		node := &ListNode{
			Val: val,
		}

		if listNode == nil {
			listNode = node
		} else {
			node.Next = listNode
			listNode = node
		}

	}

	return listNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//sum := CalNumListNode(l1) + CalNumListNode(l2)

	return nil
}

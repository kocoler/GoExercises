package lcof

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针
// a + b + c
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	hA := headA
	hB := headB
	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}

		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}

	return hA
}

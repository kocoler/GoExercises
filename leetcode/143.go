// 重排链表
// 先找中点
// 再翻转后面的
// 再合并

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    f := head
    s := head
    for f != nil && f.Next != nil && f.Next.Next != nil {
        s = s.Next
        f = f.Next.Next
    }

    mid := s
    s = s.Next
    mid.Next = nil
    tail := s
    // nx := s.Next
    for s != nil && s.Next != nil {
        s.Next, s.Next.Next, tail = s.Next.Next, tail, s.Next
    }
    // s.Next = nil

    h := head
    for h != nil && tail != nil {
        nx := h.Next
        tnx := tail.Next
        tail.Next = h.Next
        h.Next = tail
        h = nx
        tail = tnx
    }
}
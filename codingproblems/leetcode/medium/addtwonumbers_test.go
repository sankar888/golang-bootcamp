package medium

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
https://leetcode.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func next(node *ListNode) (next *ListNode, end bool) {
	if n := node.Next; n != nil {
		return n, false
	}
	return nil, true
}

func (list *ListNode) String() string {
	var str []string = make([]string, 0)
	for n, end := list, false; !end; n, end = next(n) {
		str = append(str, strconv.Itoa(n.Val))
	}
	return strings.Join(str, "")
}

func createInput(number uint) *ListNode {
	var list *ListNode = nil
	var pn *ListNode = nil
	var str string = fmt.Sprintf("%d", number)
	for _, c := range str {
		in, _ := strconv.Atoi(string(c))
		node := &ListNode{
			Val:  in,
			Next: nil,
		}
		if list == nil {
			list = node
			pn = node
		} else {
			pn.Next = node
			pn = node
		}
	}
	return list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var output, cn *ListNode = nil, nil
	var cf int = 0
	n1, end1 := l1, false
	n2, end2 := l2, false
	for {
		var v1, v2 int = 0, 0
		if end1 && end2 { //break if the both list is traversed
			break
		}
		if !end1 { //each list can have different different length
			v1 = n1.Val
			n1, end1 = next(n1) //pointers havebeen moced, do not use this
		}
		if !end2 {
			v2 = n2.Val
			n2, end2 = next(n2)
		}

		//core logic
		sum := v1 + v2 + cf
		val := sum % 10
		cf = sum / 10

		//make node chain
		node := &ListNode{val, nil}
		if output == nil {
			output = node
			cn = node
		} else {
			cn.Next = node
			cn = node
		}
	}
	if cf > 0 {
		node := &ListNode{cf, nil}
		cn.Next = node
		cn = node
	}
	return output
}

func addTwoNumbersSol1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	r := result
	rl1 := l1
	rl2 := l2
	remain := 0
	for {
		total := rl1.Val + rl2.Val + remain
		r.Val, remain = total%10, total/10
		if rl1.Next == nil && rl2.Next == nil {
			if remain > 0 {
				r.Next = &ListNode{Val: remain}
			}
			return result
		}
		if rl1.Next != nil {
			rl1 = rl1.Next
		} else {
			rl1.Val = 0
		}

		if rl2.Next != nil {
			rl2 = rl2.Next
		} else {
			rl2.Val = 0
		}
		r.Next = &ListNode{}
		r = r.Next
	}
}

func TestCreateInput(t *testing.T) {
	var l1 *ListNode = createInput(9999999)
	t.Log(l1.String())
}

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode = createInput(9999999)
	var l2 *ListNode = createInput(9999)
	res := addTwoNumbers(l1, l2)
	t.Log(res.String())
}

func TestAddTwoNumbersSol1(t *testing.T) {
	var l1 *ListNode = createInput(9999999)
	var l2 *ListNode = createInput(9999)
	res := addTwoNumbersSol1(l1, l2)
	t.Log(res.String())
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	var l1 *ListNode = createInput(9999999)
	var l2 *ListNode = createInput(9999)
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}

func BenchmarkAddTwoNumbersSol1(b *testing.B) {
	var l1 *ListNode = createInput(9999999)
	var l2 *ListNode = createInput(9999)
	for i := 0; i < b.N; i++ {
		addTwoNumbersSol1(l1, l2)
	}
}

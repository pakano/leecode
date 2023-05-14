package addtwo

import "fmt"

type Node struct {
	val  int
	next *Node
}

func AddTwo(L1, L2 *Node) *Node {
	var dummyHead = &Node{val: -1}
	var p *Node = dummyHead
	var carried int
	for L1 != nil || L2 != nil {
		var val1, val2 int
		if L1 != nil {
			val1 = L1.val
			L1 = L1.next
		}
		if L2 != nil {
			val2 = L2.val
			L2 = L2.next
		}

		p.next = &Node{val: (val1 + val2 + carried) % 10}
		carried = (val1 + val2 + carried) / 10
		p = p.next
	}

	if carried != 0 {
		p.next = &Node{val: carried}
	}
	return dummyHead.next
}

func showlist(n *Node) {
	for n != nil {
		fmt.Print(n.val)
		if n.next != nil {
			fmt.Print(" -> ")
		}
		n = n.next
	}
	fmt.Println()
}

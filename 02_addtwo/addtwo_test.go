package addtwo

import "testing"

func TestAddTwo(t *testing.T) {
	n1 := Node{val: 9}
	n2 := Node{val: 9}
	n3 := Node{val: 9}
	n4 := Node{val: 9}
	n5 := Node{val: 9}
	n6 := Node{val: 9}

	n1.next = &n2
	n2.next = &n3

	n4.next = &n5
	n5.next = &n6

	showlist(AddTwo(&n1, &n4))
}

package main

import "fmt"

type Node struct {
	content interface{}
	next    *Node
	prev    *Node
}

func (n *Node) assigncontent(value interface{}) {
	n.content = value
}
func (n *Node) editenext(ptr *Node) {
	n.next = ptr
}
func (n *Node) editeprev(ptr *Node) {
	n.prev = ptr
}
type list struct {
	Head *Node
	Tail *Node
}
func (l *list) addfornt(value interface{}) {
	var newnode Node
	newnode.assigncontent(value)
	//created the new node
	if l.Head != nil {
		l.Head.prev = &newnode
		newnode.editenext(l.Head)
	} else {
		l.Tail = &newnode
	}
	//cheaked if the head is nil or not
	//then assigning the old headnode's prev to be  the new node (the head ) if the head wasn't nil
	//then assigning the head to the new nod if head
	l.Head = &newnode

}
func (l *list) AddToTail(value interface{}) {
	var newnode Node
	newnode.assigncontent(value)
	//created the new node
	if l.Tail != nil {
		newnode.editeprev(l.Tail)
		//point the prev of newnode to old tail
		l.Tail.next = &newnode
		// point the old tail's next to the new node (the new tail )
		l.Tail = &newnode
		//point the list tail to the new tail (new node )
	} else {
		//if the list is empty the prev and the next pointers both will point the new  node .
		l.Tail = &newnode
		l.Head = &newnode
	}

}
func (l *list) AddAt(idx int, value interface{}) {
	holder := l.Head
	for i := 1; i < idx; {
		if holder.next != nil {
			holder = holder.next
			i++
		} else {
			fmt.Println("out of index")
			break // should handle this later
		}

	}
	newnode:= &Node{}
	newnode.assigncontent(value)
	newnode.editeprev(holder)
	newnode.editenext(holder.next)
	holder.editenext(newnode)


}
func (l *list) GetNodeAt(idx int)*Node{
	holder := l.Head
	for i := 1; i < idx; {
		if holder.next != nil {
			holder = holder.next
			i++
		} else {
			fmt.Println("out of index")
			break // should handle this later
		}
	}
	return holder
}
func (l *list)delete(idx int ){
	holder := l.Head
	for i := 1; i < idx; {
		if holder.next != nil {
			holder = holder.next
			i++
		} else {
			fmt.Println("out of index")
			break // should handle this later
		}
	}
	prev := holder.prev
	next := holder.next
	prev.next = next
	next.prev = prev
}
func main() {
	llist := list{}
	llist.addfornt(3)
	llist.addfornt(467)
	llist.AddToTail(20)
	llist.AddAt(2,700)
	fmt.Println(llist.GetNodeAt(3).content)
	llist.delete(3)
	fmt.Println(llist.GetNodeAt(3).content)

}

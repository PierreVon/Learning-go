package main

import (
	"errors"
	"fmt"
)

// 节点
type node struct {
	val interface{}
	next *node
}

// 接收器，即方法，定义在结构体外
func (n *node) setVal(val interface{}){
	n.val = val
}

func (n *node) getVal() interface{}{
	return n.val
}

func (n *node) setNext(next *node) {
	n.next = next
}

func (n *node) getNext() *node{
	return n.next
}

// 列表
type list struct {
	head node
}

// 列表接收器
func (l *list) insert(n *node) {
	p := &l.head
	for p.next != nil{
		p = p.next
	}
	p.next = n
}

func (l *list) delete(val interface{}) (err error) {
	p := &l.head
	pre := &l.head
	if p.next == nil {
		err = errors.New("list is empty")
	}else{
		p = p.next
	}
	for p != nil{
		if p.val == val {
			pre.next = p.next
			return nil
		}
		p = p.next
		pre = pre.next
	}
	err = errors.New("no such node")
	return
}

func (l *list) display() {
	p := &l.head
	for p.next != nil{
		p = p.next
		fmt.Println(p.val)
	}
}



func main() {
	head := node{
		-1,
		nil,
	}
	l := list{head}
	// vals := []int{1,2,3,4}
	vals := []string{"lucy", "mike", "pierre"}
	for _, val := range vals{
		n := node{
			val,
			nil,
		}
		l.insert(&n)
	}
	l.display()
	l.delete("mike")
	l.display()

}


package main

import (
	"container/list"
	"fmt"
	"sync"
	"errors"
)

func listOperation() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

//реализация стека
type stackWM []int

func (s stackWM) Push(v int) stackWM {
	return append(s, v)
}

func (s stackWM) Pop() (stackWM, int) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return  s[:l-1], s[l-1]
}


//потоко безопасный метод
type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s []int
}

func NewStack() *stack {
	return &stack {sync.Mutex{}, make([]int,0), }
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()


	l := len(s.s)
	if l == 0 {
		return 0, errors.New("стек пуст")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func main()  {
	listOperation()
	fmt.Println("------")
	s := NewStackList()
	s.Push(7)
	s.Push(8)
	s.Push(9)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

//реализация LIFO с использованием структуры связанных данных
type element struct {
	data interface{}
	next *element
}

type stackList struct {
	lock *sync.Mutex
	head *element
	Size int
}

func (stk *stackList) Push(data interface{}) {
	stk.lock.Lock()

	element := new(element)
	element.data = data
	temp := stk.head
	element.next = temp
	stk.head = element
	stk.Size++

	stk.lock.Unlock()
}

func (stk *stackList) Pop() interface{} {
	if stk.head == nil {
		return nil
	}
	stk.lock.Lock()
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	stk.lock.Unlock()

	return r
}

func NewStackList() *stackList {
	stk := new(stackList)
	stk.lock = &sync.Mutex{}

	return stk
}
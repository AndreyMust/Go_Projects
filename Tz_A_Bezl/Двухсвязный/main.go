package main

import (
	"errors"
	"fmt"
)

// Интерфейс с требуемыми методами
type MyInterface[T any] interface {
	Delete(func(item T) bool) bool
	Reverse()
	Swap(i, j int) error
}

// Node представляет узел двусвязного списка
type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

// DoublyLinkedList представляет двусвязный список
type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewDoublyLinkedList создает новый пустой двусвязный список
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Add добавляет новый элемент в конец списка
func (dll *DoublyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}
	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.size++
}

// Print печатает элементы списка
func (dll *DoublyLinkedList[T]) Print() {
	current := dll.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

// Swap меняет местами элементы с индексами i и j
func (dll *DoublyLinkedList[T]) Swap(i, j int) error {
	if i >= dll.size || j >= dll.size {
		return errors.New("индекс вне границ списка")
	}

	if i == j {
		return nil // Ничего менять не нужно
	}

	var nodeI, nodeJ *Node[T]
	current := dll.head

	for idx := 0; current != nil && (nodeI == nil || nodeJ == nil); idx++ {
		if idx == i {
			nodeI = current
		}
		if idx == j {
			nodeJ = current
		}
		current = current.next
	}
	if nodeI != nil && nodeJ != nil {
		nodeI.value, nodeJ.value = nodeJ.value, nodeI.value
		return nil
	}
	return errors.New("не удалось найти указанные индексы")
}

// Reverse разворачивает двусвязный список
func (dll *DoublyLinkedList[T]) Reverse() {
	current := dll.head
	var prev *Node[T] = nil
	dll.tail = dll.head

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	dll.head = prev
}

func main() {
	fmt.Println("Hello")

	list := NewDoublyLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()
	list.Reverse()
	list.Print()
	list.Swap(0, 1)
	list.Print()

}

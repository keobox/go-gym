package main

import "fmt"

// Generic function that returns the keys of a map
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Generic linked list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Methods
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.head.next
	}

}

func (lst *List[T]) GetAll() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func main() {

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// MapKeys can infer the key and value types
	fmt.Println("keys:", MapKeys(m))

	// but is possible to be explicit
	_ = MapKeys[int, string](m)

	// use the single linked list
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

}

package main

import (
	"fmt"
	"learning-go/structures/linkedList"
	"learning-go/structures/queue"
	"math/rand"
)

func main() {
	// Linked List
	list := linkedlist.LinkedList[int]{}
	for i := 0; i < 100; i++ {
		list.Add(rand.Int())
	}
	for i := range list.Iter() {
		fmt.Println(i)
	}

	// Queue
	q := queue.Queue[int]{}
	for i := 0; i < 200; i++ {
		q.Enqueue(rand.Int())
	}
	for i := 0; i < q.Len(); i++ {
		fmt.Println("q:", q.At(i))
	}
}

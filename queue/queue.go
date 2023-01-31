package queue

import (
	"fmt"

	"github.com/nikhilnarayanan623/stackAndQueue/queue/interfaces"
)

// ro hold single value and connect to next node
type node struct {
	data int
	next *node
}

// to hold the whole queue and its refrence function
type queue struct {
	head *node
	tail *node
}

// this function return a new queueu instance
func NewQueue() interfaces.QueueInterface {
	return &queue{}
}

// function to insert a new value to the end of queue
func (q *queue) Enqueue(data int) {

	newNode := &node{data: data}
	if checkQueueIsEmpty(q.head, false) { //head and the bool is for need to show the message or not
		q.head = newNode
	} else {
		q.tail.next = newNode
	}
	q.tail = newNode
}

// function to get multiple values form the user and add to queue
func (q *queue) MultipleEnqueue() {

	var limit, data int

	fmt.Print("Enter how much value you need add to queue: ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		fmt.Print("Enter value: ")
		fmt.Scan(&data)
		q.Enqueue(data)
	}
}

// function to delete a value from the begining of the queue
func (q *queue) Dequeue() {

	if checkQueueIsEmpty(q.head, true) {
		return
	}
	// if the queue have values then move head to its next
	//so the value in begining has been detached from queue
	q.head = q.head.next
}

// function to display all values in the queue
func (q *queue) DisplayOrder() {

	if checkQueueIsEmpty(q.head, true) {
		return
	}
	fmt.Println("\nValues in the Queue")
	currentNode := q.head

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

// functoin to check a the queue is nil or not
func checkQueueIsEmpty(node *node, show bool) bool {
	//if node is nill and show is true then show the message
	if node == nil && show {
		fmt.Println("Queue is empty")
	}
	//return node is nil or not
	return node == nil
}

package queue

import (
	"errors"
	"log"
	"os"
	"reflect"
)

//Element - struct that describes element in the Queue
//next - field that's describes the reference on the next Element
//value - field that's describes stored value in the Queue
type Element struct {
	value interface{}
	next  *Element
}

//Queue - struct that implements Singly Queue data structures
//head - reference of the Queue head
//tail - reference of the Queue tail
//size - the current number of elements in Queue
//cap - the max capacity of the Queue
type Queue struct {
	head *Element
	tail *Element
	size int
	cap  int
}

//New - return a new Queue instance
//maxSize - the Queue's capacity value
func New(maxSize int) *Queue {
	return &Queue{cap: maxSize}
}

//GetSize - return Queue's size
func (q *Queue) GetSize() int {
	return q.size
}

//GetCap - return Queue's capacity
func (q *Queue) GetCap() int {
	return q.cap
}

//typeDefinition - defines the actual type of the Queue's elements
func typeDefinition(i1 interface{}, i2 interface{}) (bool, error) {
	if reflect.TypeOf(i1) != reflect.TypeOf(i2) {
		err := errors.New("type incompatibility ")
		return false, err
	}
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int), nil
	case float64:
		return i1.(float64) > i2.(float64), nil
	case string:
		return i1.(string) > i2.(string), nil
	case rune:
		return i1.(rune) > i2.(rune), nil
	case uint:
		return i1.(uint) > i2.(uint), nil
	case byte:
		return i1.(byte) > i2.(byte), nil
	default:
		err := errors.New("this type isn't supported")
		return false, err
	}
}

//Sort - sorts Queue using bumble sort algorithm
func (q *Queue) Sort() {
	cur := q.head
	for cur != nil {
		nextElement := cur.next
		for nextElement != nil {
			state, err := typeDefinition(cur.value, nextElement.value)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			} else {
				if state {
					temp := cur.value
					cur.value = nextElement.value
					nextElement.value = temp
				}
				nextElement = nextElement.next
			}
		}
		cur = cur.next
	}
}

//IsFull - Check if the Queue is full
func (q *Queue) IsFull() bool {
	return q.cap <= q.size
}

//IsEmpty - Check if the Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

//Peek - Get the value of the front of the Queue without removing it. Return an error if Queue is empty
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty. You can't display element")
	}
	return q.head.value, nil

}

//Enqueue - Add an element to the end of the Queue. Return an error if Queue is full
//value - element that adds to the end of the Queue, the type is interface{}
func (q *Queue) Enqueue(value interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full. You can't add element in your queue")
	}
	element := &Element{
		value: value,
	}
	if q.head == nil {
		q.tail = element
		q.head = element
	} else {
		q.tail.next = element
		q.tail = element
	}
	q.size++
	return nil
}

//Dequeue - Remove an element from the front of the Queue. Return an error if Queue is empty
func (q *Queue) Dequeue() error {
	if q.head == nil {
		return errors.New("your queue is empty and you can't delete element")
	}
	element := Element{}
	if q.head != nil {
		element.value = q.head.value
		if q.head.next != nil {
			q.head = q.head.next
		} else {
			q.head = nil
			q.tail = nil
		}
	}
	q.size--
	return nil
}

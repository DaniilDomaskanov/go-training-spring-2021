package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_GetCap(t *testing.T) {
	assert := assert.New(t)
	actual := New(6).GetCap()
	expected := 6
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected", expected, actual))
}

func TestQueue_GetSize(t *testing.T) {
	assert := assert.New(t)
	queue := New(6)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	actual := queue.GetSize()
	expected := 3
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected", expected, actual))
}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	actual := New(5)
	expected := &Queue{
		size: 0,
		cap:  5,
	}
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", expected, actual))
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	queue := New(3)
	actual := queue.IsEmpty()
	expected := true
	assert.Equal(expected, actual, fmt.Sprintf("%t and %t not equal, but expected", expected, actual))
}

func TestQueue_IsEmpty2(t *testing.T) {
	assert := assert.New(t)
	queue := New(3)
	queue.Enqueue(1)
	actual := queue.IsEmpty()
	expected := false
	assert.Equal(expected, actual, fmt.Sprintf("%t and %t not equal, but expected", expected, actual))
}

func TestQueue_IsFull2(t *testing.T) {
	assert := assert.New(t)
	queue := New(3)
	queue.Enqueue("a")
	queue.Enqueue("b")
	actual := queue.IsFull()
	expected := false
	assert.Equal(expected, actual, fmt.Sprintf("%t and %t not equal, but expected", expected, actual))
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	queue := New(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	actual := queue.IsFull()
	expected := true
	assert.Equal(expected, actual, fmt.Sprintf("%t and %t not equal, but expected", expected, actual))
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	queue := New(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	actual, _ := queue.Peek()
	expected := 1
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected", expected, actual))
}

func TestQueue_Peek2(t *testing.T) {
	assert := assert.New(t)
	queue := New(2)
	_, err := queue.Peek()
	assert.EqualError(err, "queue is empty. You can't display element")
}

func TestQueue_Peek3(t *testing.T) {
	assert := assert.New(t)
	queue := New(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Peek()
	actual := queue.GetSize()
	expected := 2
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected", expected, actual))
}

func TestQueue_Sort(t *testing.T) {
	assert := assert.New(t)
	actual := New(5)
	actual.Enqueue(2)
	actual.Enqueue(3)
	actual.Enqueue(10)
	actual.Enqueue(7)
	actual.Enqueue(0)
	actual.Sort()
	expected := []int{0, 2, 3, 7, 10}
	for i := 0; i < actual.GetSize(); i++ {
		current := actual.head
		for j := 0; j < i; j++ {
			current = current.next
		}
		actualVal := current.value
		expected := expected[i]
		assert.Equal(expected, actualVal, fmt.Sprintf("%v and %v not equal, but expected", expected, actualVal))
	}
}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	actual := New(5)
	err := actual.Dequeue()
	assert.EqualError(err, "your queue is empty and you can't delete element")
}

func TestQueue_Dequeue2(t *testing.T) {
	assert := assert.New(t)
	queue := New(5)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	actual, _ := queue.Peek()
	expected := 3
	assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", expected, actual))
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	actual := New(0)
	err := actual.Enqueue(1)
	assert.EqualError(err, "queue is full. You can't add element in your queue")
}

func TestQueue_Enqueue2(t *testing.T) {
	assert := assert.New(t)
	queue := New(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	expected := []int{1, 2, 3}
	for i := 0; i < queue.GetSize(); i++ {
		curr := queue.head
		for j := 0; j < i; j++ {
			curr = curr.next
		}
		actual := curr.value
		assert.Equal(expected[i], actual, fmt.Sprintf("%v and %v not equal, but expected", expected, actual))
	}
}

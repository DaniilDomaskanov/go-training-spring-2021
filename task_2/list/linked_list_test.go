package list

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_GetLen(t *testing.T) {
	assert := assert.New(t)
	actual := New(1, 2, 3).GetLen()
	expected := 3
	assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected", expected, actual))
}

func TestLinkedList_Search(t *testing.T) {
	assert := assert.New(t)
	actual := New(1, 2, 3, 4, 5, 10)
	id := 9
	_, err := actual.Search(id)
	assert.EqualError(err, fmt.Sprintf("index - %d is greater than LinkedList length - %d", id, actual.GetLen()))

}

func TestLinkedList_Search2(t *testing.T) {
	assert := assert.New(t)
	actual, _ := New("a", "b", "c").Search(0)
	expected := "c"
	assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected", expected, actual))

}

func TestLinkedList_Search3(t *testing.T) {
	assert := assert.New(t)
	actual := New("a", "b", "c", "d")
	expected := []string{"d", "c", "b", "a"}
	for i := 0; i < actual.GetLen(); i++ {
		actual, _ := actual.Search(i)
		expected := expected[i]
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected", expected, actual))
	}
}

func TestLinkedList_Delete(t *testing.T) {
	assert := assert.New(t)
	actual := New("a", "b", "c")
	expected := []string{"c", "a"}
	actual.Delete(1)
	for i := 0; i < actual.GetLen(); i++ {
		actual, _ := actual.Search(i)
		expected := expected[i]
		assert.Equal(expected, actual, fmt.Sprintf("%s and %s not equal, but expected", expected, actual))
	}
}

func TestLinkedList_Delete2(t *testing.T) {
	assert := assert.New(t)
	actual := New("a", "b", "c")
	id := -1
	err := actual.Delete(id)
	assert.EqualError(err, fmt.Sprintf("id is a negative value - %d", id))
}

func TestLinkedList_Delete3(t *testing.T) {
	assert := assert.New(t)
	actual := New()
	err := actual.Delete(2)
	assert.EqualError(err, "delete operation isn't support because list is empty")
}

func TestLinkedList_Deletion2(t *testing.T) {
	assert := assert.New(t)
	actual := New(1, 2, 3)
	expected := []int{2, 1}
	actual.Deletion()
	for i := 0; i < actual.GetLen(); i++ {
		actual, _ := actual.Search(i)
		expected := expected[i]
		assert.Equal(expected, actual, fmt.Sprintf("%d and %d not equal, but expected", expected, actual))
	}

}

func TestLinkedList_Deletion(t *testing.T) {
	assert := assert.New(t)
	actual := &LinkedList{}
	err := actual.Deletion()
	assert.EqualError(err, "linked list is empty")
}

func TestLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	actual := New(10, 1, 2)
	actual.Display()
	actual.Sort()
	actual.Display()
	expected := []int{1, 2, 10}
	for i := 0; i < actual.GetLen(); i++ {
		expected := expected[i]
		actual, _ := actual.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", expected, actual))
	}
}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	actual := New(0, 1, 2)
	expected := &LinkedList{}
	expected.Insert(0)
	expected.Insert(1)
	expected.Insert(2)
	for i := 0; i < actual.len; i++ {
		actual, _ := actual.Search(i)
		expected, _ := expected.Search(i)
		assert.Equal(expected, actual, fmt.Sprintf("%v and %v not equal, but expected", expected, actual))
	}
}

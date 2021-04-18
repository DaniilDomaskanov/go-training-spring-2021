package list

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

//Element - struct that describes element in the LinkedList
//next - field that's describes the reference on the next Element
//value - field that's describes stored value in the LinkedList
type Element struct {
	next  *Element
	value interface{}
}

//LinkedList - struct that implements Singly LinkedList data structures
//head - reference of the LinkedList head
//len - the count of elements in LinkedList
type LinkedList struct {
	head *Element
	len  int
}

//GetLen - returns LinkedList length
func (l *LinkedList) GetLen() int {
	return l.len
}

//increaseLen - increases LinkedList len field

func (l *LinkedList) increaseLen() {
	l.len++
}

//decreaseLen - decreases LinkedList len field
func (l *LinkedList) decreaseLen() {
	l.len--
}

//New create a new LinkedList
//values - varargs that perform LinkedList values. The type is interface{}
func New(values ...interface{}) *LinkedList {
	linkedList := &LinkedList{}
	for _, value := range values {
		linkedList.Insert(value)
	}
	return linkedList
}

//Insert - Adds an element at the beginning of the list.
//key - element that adds to the LinkedList, the type is interface{}
func (l *LinkedList) Insert(key interface{}) {
	element := &Element{
		value: key,
	}
	element.next = l.head
	l.head = element
	l.increaseLen()

}

//Deletion - Deletes an element at the beginning of the list.Return an error if LinkedList is empty
func (l *LinkedList) Deletion() error {
	if l.len == 0 {
		return errors.New("linked list is empty")
	}
	l.head = l.head.next
	l.decreaseLen()
	return nil
}

//getElementById - return an LinkedList's element according to their id. If id is incorrect returns a error
//id - element's id in the LinkedList, the type is int
func (l *LinkedList) getElementById(id int) (*Element, error) {
	if id < 0 {
		return nil, errors.New(fmt.Sprintf("id is a negative value - %d", id))
	}
	if id >= l.len {
		return nil, errors.New(fmt.Sprintf("index - %d is greater than LinkedList length - %d", id, l.len))
	}
	current := l.head
	for i := 0; i < id; i++ {
		current = current.next
	}
	return current, nil
}

//Search - Searches an element using the id.Return an error if element not found or index is incorrect otherwise return the element using the id
//id - element's id in the LinkedList, the type is int
func (l *LinkedList) Search(id int) (interface{}, error) {
	current, err := l.getElementById(id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return nil, errors.New("element not found")
	}
	return current.value, nil
}

//Delete - Deletes an element using the id. Return an error if LinkedList is empty
//id - element's id in the LinkedList, the type is int
func (l *LinkedList) Delete(id int) error {
	if l.len == 0 {
		return errors.New("delete operation isn't support because list is empty")
	}
	curr, err := l.getElementById(id)
	if err != nil {
		return err
	}
	if id == 0 {
		return l.Deletion()
	} else {
		prev, err := l.getElementById(id - 1)
		if err != nil {
			return err
		}
		if curr == l.head {
			l.head = l.head.next
		} else {
			prev.next = curr.next
		}
		l.decreaseLen()
	}
	return nil
}

//Display - Displays the complete list.
func (l *LinkedList) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%v ", list.value)
		list = list.next
	}
	fmt.Println()
}

//typeDefinition - defines the actual type of the LinkedList's elements
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

//Sort - sorts LinkedList using bumble sort algorithm
func (l *LinkedList) Sort() {
	cur := l.head
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

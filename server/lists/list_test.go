package lists

import (
	"reflect"
	"testing"
	"time"
)

func TestNewBlankTodoList(t *testing.T) {
	title := "Test List"
	list := NewBlankTodoList(title)

	if list.Title != title {
		t.Errorf("Expected list.Title to be %s, but got %s", title, list.Title)
	}

	if len(list.Items) != 0 {
		t.Errorf("Expected list.Items length to be 0, but got %v", list.Items)
	}

	if reflect.TypeOf(list.Items).Kind().String() != "slice" {
		t.Errorf("Expected list.Items type to be slice, but got %s", reflect.TypeOf(list.Items).Kind().String())
	}

	if reflect.TypeOf(list.Items).Elem().Name() != "TodoItem" {
		t.Errorf("Expected list.Items element type to be TodoItem, but got %s", reflect.TypeOf(list.Items).Elem().Name())
	}
}

func TestNewTodoList(t *testing.T) {
	addedOn := time.Now()
	title := "Test List"
	list := NewTodoList(title, []TodoItem{
		TodoItem{"Foo", addedOn},
	})

	if len(list.Items) != 1 {
		t.Errorf("Expected list.Items length to be 1, but got %d", len(list.Items))
	}

	if list.Items[0].Desc != "Foo" {
		t.Errorf(`Expected list.Items[0].Desc to be "Foo", but got %s`, list.Items[0].Desc)
	}

	if list.Items[0].AddedOn != addedOn {
		t.Errorf(`Expected list.Items[0].AddedOn to be time.Now(), but got %v`, list.Items[0].AddedOn)
	}
}

func TestTodoListAddItem(t *testing.T) {
	title := "Test List"
	list := NewBlankTodoList(title)
	list.AddItem("Something todo")

	if len(list.Items) != 1 {
		t.Errorf("Expected list.Items length to be 1, but got %d", len(list.Items))
	}
}

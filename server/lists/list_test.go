package lists

import (
	"reflect"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
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
		TodoItem{
			ID:      uuid.NewV4().String(),
			Desc:    "Foo",
			AddedOn: addedOn,
		},
	})

	if len(list.Items) != 1 {
		t.Errorf("Expected list.Items length to be 1, but got %d", len(list.Items))
	}

	if list.Items[0].Desc != "Foo" {
		t.Errorf(`Expected list.Items[0].Desc to be "Foo", but got %s`, list.Items[0].Desc)
	}

	if list.Items[0].AddedOn != addedOn {
		t.Errorf("Expected list.Items[0].AddedOn to be time.Now(), but got %v", list.Items[0].AddedOn)
	}
}

func TestTodoListIDShouldNotBeEmpty(t *testing.T) {
	list1 := NewBlankTodoList("Test List 1")

	if len(list1.ID) == 0 {
		t.Errorf("Expected list1.ID to be not empty")
	}
}

func TestTodoListIDShouldBeUnique(t *testing.T) {
	list1 := NewBlankTodoList("Test List 1")
	list2 := NewBlankTodoList("Test List 2")

	if list1.ID == list2.ID {
		t.Errorf("Expected list1.ID to be not equal to list2.ID, but got same values: %s", list1.ID)
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

func TestTodoListAddItemShouldReturnID(t *testing.T) {
	list := NewBlankTodoList("Another list")
	itemID := list.AddItem("Something todo")

	if len(itemID) == 0 {
		t.Error("Expected list.AddItem to return ID")
	}
}

func TestTodoListRemoveItem(t *testing.T) {
	list := NewBlankTodoList("Another list")
	itemID := list.AddItem("Something todo")
	list.AddItem("Another thing todo")
	count, err := list.RemoveItem(itemID)

	if count != 1 {
		t.Errorf("Expected count returned from list.RemoveItem to be 1, but got %d instead", count)
	}

	if err != nil {
		t.Error("Expected list.RemoveItem to not return error")
	}
}

func TestTodoListRemoveItemShouldReturnError(t *testing.T) {
	list := NewBlankTodoList("Another list")
	list.AddItem("Something todo")
	_, err := list.RemoveItem("fake item ID")

	if err == nil {
		t.Error("Expected list.RemoveItem to return error")
	}

	if err.Error() != "Item not found" {
		t.Error(`Expected list.RemoveItem to return "Item not found" error`)
	}
}

func TestTodoListAddItemIDShouldBeUnique(t *testing.T) {
	title := "Test List"
	list := NewBlankTodoList(title)
	list.AddItem("Something todo")
	list.AddItem("Something todo")

	if list.Items[0].ID == list.Items[1].ID {
		t.Errorf("Expected list.Items IDs to be unique")
	}
}

func TestTodoItemShouldHaveID(t *testing.T) {
	list := NewBlankTodoList("Test List")
	list.AddItem("Something todo")

	if len(list.Items[0].ID) == 0 {
		t.Errorf("Expected list.Items[0].ID to be not empty")
	}
}

func TestNewTodoItem(t *testing.T) {
	item := NewTodoItem("Another thing todo")

	if len(item.ID) == 0 {
		t.Errorf("Expected item.ID to be not empty")
	}

	if len(item.Desc) == 0 {
		t.Errorf("Expected item.Desc to be not empty")
	}
}

func TestNewTodoBoard(t *testing.T) {
	board := NewTodoBoard()

	if len(board.ID) == 0 {
		t.Errorf("Expected board.ID to be not empty")
	}

	if reflect.TypeOf(board.Lists).Elem().Name() != "TodoList" {
		t.Errorf("Expected board.Lists to be slice of TodoList")
	}
}

package lists

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TodoItem represents a item in todo list
type TodoItem struct {
	ID      string
	Desc    string
	AddedOn time.Time
}

// TodoList represents a todo list
type TodoList struct {
	ID    string
	Title string
	Items []TodoItem
}

// TodoBoard manages a group of TodoLists
type TodoBoard struct {
	ID    string
	Lists []TodoList
}

// NewBlankTodoList returns new list object
func NewBlankTodoList(title string) TodoList {
	list := TodoList{
		ID:    uuid.NewV4().String(),
		Title: title,
		Items: make([]TodoItem, 0),
	}
	return list
}

// NewTodoList returns new list object with predefined todo items
func NewTodoList(title string, items []TodoItem) TodoList {
	list := TodoList{
		ID:    uuid.NewV4().String(),
		Title: title,
		Items: items,
	}
	return list
}

// NewTodoItem returns new item that can be added to todo list
func NewTodoItem(desc string) TodoItem {
	item := TodoItem{
		ID:      uuid.NewV4().String(),
		Desc:    desc,
		AddedOn: time.Now(),
	}
	return item
}

// NewTodoBoard returns new TodoBoard
func NewTodoBoard() TodoBoard {
	board := TodoBoard{
		ID: uuid.NewV4().String(),
	}
	return board
}

// AddItem adds new todo item to TodoList
func (list *TodoList) AddItem(desc string) {
	item := NewTodoItem(desc)
	list.Items = append(list.Items, item)
}

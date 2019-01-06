package lists

import "time"

// TodoItem represents a item in todo list
type TodoItem struct {
	Desc    string
	AddedOn time.Time
}

// TodoList represents a todo list
type TodoList struct {
	Title string
	Items []TodoItem
}

// AddItem adds new todo item to TodoList
func (list *TodoList) AddItem(desc string) {
	item := TodoItem{
		Desc:    desc,
		AddedOn: time.Now(),
	}
	list.Items = append(list.Items, item)
}

// NewBlankTodoList returns new list object
func NewBlankTodoList(title string) TodoList {
	list := TodoList{
		Title: title,
		Items: make([]TodoItem, 0),
	}
	return list
}

// NewTodoList returns new list object with predefined todo items
func NewTodoList(title string, items []TodoItem) TodoList {
	list := TodoList{
		Title: title,
		Items: items,
	}
	return list
}

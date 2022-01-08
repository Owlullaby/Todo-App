package models

type Todo struct {
	ID int64
	Desc string
}

type TodoList struct {
	Todos []Todo
}
package model

type Task struct {
	TaskName   string
	TaskId     string
	TaskDetail string
	TaskTypeN  string
}

type TaskType struct {
	TaskTypeId   string
	TaskTypeName string
}

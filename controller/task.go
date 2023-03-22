package controller

import (
	"fmt"
	"test/model"
)

type Task struct {
	Id   int
	Text string
}

func NewTask() Task {
	return Task{}
}

func (c Task) Create(text string) {
	var repo model.TaskRepository = model.NewTaskRepository()
	repo.Create(text)
}

func (c Task) GetAll() interface{} {
	var repo model.TaskRepository = model.NewTaskRepository()
	var tasks model.Tasks = repo.GetAll()
	return tasks
}

func (c Task) GetById(id int) interface{} {
	var repo model.TaskRepository = model.NewTaskRepository()
	var tasks model.Tasks = repo.GetById(id)
	fmt.Println("cont:", tasks)
	return tasks
}

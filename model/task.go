package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func init() {
}

type Task struct {
	Id   int    `gorm:"primary_key"`
	Text string `gorm:"size:140"`
}
type Tasks []Task

type TaskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return TaskRepository{}
}

func (m TaskRepository) Create(text string) {
	var task Task = Task{Text: text}
	//db.NewRecord(task)
	db.Create(&task)
	db.Save(&task)
}

func (m TaskRepository) GetAll() Tasks {
	var tasks Tasks = Tasks{}
	db.Find(&tasks)
	return tasks
}

func (m TaskRepository) GetById(id int) Tasks {
	var tasks Tasks = Tasks{}
	db.Find(&tasks, id)
	fmt.Println("model", tasks)
	return tasks
}

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "task.db")
	db.DropTableIfExists(&Task{})
	db.CreateTable(&Task{})

	if err != nil {
		panic(err)
	}
}

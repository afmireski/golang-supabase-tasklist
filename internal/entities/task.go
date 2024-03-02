package entities

import (
	"time"
)

type Task struct {
	Id          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Finished    bool      `json:"completed"`
	Creator     *Creator  `json:"creator"`
}

func NewTask(id int32, title string, description string, date time.Time, finished bool, creator *Creator) *Task {
	return &Task{
		Id:          id,
		Title:       title,
		Description: description,
		Date:        date,
		Finished:    finished,
		Creator:     creator,
	}
}

package entities

import "time"

type Task struct {
	Id          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Finished    bool      `json:"completed"`
	Creator     Creator  `json:"creator"`
}

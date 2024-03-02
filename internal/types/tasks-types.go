package types

import "time"

type CreateTaskInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Finished    bool      `json:"completed"`
	CreatorId   string    `json:"creatorId"`
}
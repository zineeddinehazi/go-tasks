package pkg

import (
	"database/sql"
)

type Task struct {
	ID      string `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	IsDone  bool   `json:"isdone" db:"isdone"`
}

type TaskRepository struct {
	DB *sql.DB
}

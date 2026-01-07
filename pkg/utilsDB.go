package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (repo *TaskRepository) CreateTable() error {
	_, err := repo.DB.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT NOT NULL,
        isdone BOOLEAN
    )`)
	return err
}

func (repo *TaskRepository) QuerryList() ([]Task, error) {
	rows, err := repo.DB.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Content, &t.IsDone); err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, rows.Err()
}

func (repo *TaskRepository) InsertTask(task Task) error {
	_, err := repo.DB.Exec(`INSERT INTO TASKS (content, isdone) VALUES (?, ?)`, task.Content, task.IsDone)
	return err
}

func (repo *TaskRepository) QuerryTask(ID string) (any, error) {
	row := repo.DB.QueryRow(`SELECT id, content, isdone FROM tasks WHERE id=?`, ID)
	var t Task
	if err := row.Scan(&t.ID, &t.Content, &t.IsDone); err != nil {
		if err == sql.ErrNoRows {
			return map[string]string{"message": fmt.Sprintf("No task found with id %s", ID)}, nil
		}
		return nil, err
	}
	return &t, nil
}

func (repo *TaskRepository) UpdateTask(ID string) error {
	result, err := repo.DB.Exec(`
        UPDATE tasks SET isdone = (isdone + 1) % 2  WHERE id = ?`, ID)
	handleErr(err)

	rows, err := result.RowsAffected()
	handleErr(err)
	if rows == 0 {
		return fmt.Errorf("No task found with id %s", ID)
	}
	return nil
}

func (repo *TaskRepository) DeleteTask(ID string) error {
	result, err := repo.DB.Exec("DELETE FROM tasks WHERE id = ?", ID)
	handleErr(err)

	rows, err := result.RowsAffected()
	handleErr(err)
	if rows == 0 {
		return fmt.Errorf("No task found with id %s", ID)
	}
	return nil
}

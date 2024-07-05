package postgres

import (
	"Tasks/internal/models"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type Task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *Task {
	return &Task{db: db}
}

func (t *Task) StoreCreateTask(task *models.Task) (*models.Task, error) {
	var tasks models.Task
	sql, args, err := squirrel.
		Insert("tasks").
		Columns("title", "description").
		Values(task.Title, task.Description).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to insert task: %v", err)
	}

	row := t.db.QueryRow(sql, args...)
	if err := row.Scan(&tasks.ID); err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}

	return &tasks, nil
}

func (u *Task) StoreGetbyIdTasks(id int32) (*models.Task, error) {
	var task models.Task
	sql, args, err := squirrel.
		Select("*").
		From("tasks").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to select task: %v", err)
	}
	row := u.db.QueryRow(sql, args...)
	if err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
		return nil, fmt.Errorf("to scan error: %v", err)
	}

	return &task, nil
}

func (u *Task) StoreGetTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	sql, args, err := squirrel.
		Select("*").
		From("tasks").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to select task: %v", err)
	}
	rows, err := u.db.Query(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t *Task) StoreUpdateTask(task *models.Task) (*models.Task, error) {
	sql, args, err := squirrel.
		Update("tasks").
		Set("title", task.Title).
		Set("description", task.Description).
		Set("done", task.Done).
		Where(squirrel.Eq{"id": task.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to update task: %v", err)
	}

	_, err = t.db.Exec(sql, args...)
	if err != nil {
		return nil, fmt.Errorf("update error: %v", err)
	}

	return task, nil
}

func (t *Task) StoreDeleteTask(id int32) error {
	sql, args, err := squirrel.
		Delete("tasks").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("unable to delete task: %v", err)
	}

	_, err = t.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("delete error: %v", err)
	}

	return nil
}

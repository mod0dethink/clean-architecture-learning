package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"

	"clean-architecture-learning/backend/internal/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Migrate はアプリ起動時に一度だけ呼び出してテーブルを作成する。
func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id         TEXT    PRIMARY KEY,
			title      TEXT    NOT NULL,
			status     TEXT    NOT NULL,
			created_at INTEGER NOT NULL
		)`)
	return err
}

func (r *TaskRepository) Save(task domain.Task) error {
	_, err := r.db.Exec(`
		INSERT INTO tasks (id, title, status, created_at)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET status = excluded.status`,
		task.ID, task.Title, string(task.Status), task.CreatedAt.Unix())
	return err
}

func (r *TaskRepository) FindAll() ([]domain.Task, error) {
	rows, err := r.db.Query(
		`SELECT id, title, status, created_at FROM tasks ORDER BY created_at ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var t domain.Task
		var status string
		var createdAtUnix int64
		if err := rows.Scan(&t.ID, &t.Title, &status, &createdAtUnix); err != nil {
			return nil, err
		}
		t.Status = domain.Status(status)
		t.CreatedAt = time.Unix(createdAtUnix, 0)
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (r *TaskRepository) FindByID(id string) (domain.Task, error) {
	row := r.db.QueryRow(
		`SELECT id, title, status, created_at FROM tasks WHERE id = ?`, id)
	var t domain.Task
	var status string
	var createdAtUnix int64
	if err := row.Scan(&t.ID, &t.Title, &status, &createdAtUnix); err != nil {
		return domain.Task{}, fmt.Errorf("task %s not found: %w", id, err)
	}
	t.Status = domain.Status(status)
	t.CreatedAt = time.Unix(createdAtUnix, 0)
	return t, nil
}

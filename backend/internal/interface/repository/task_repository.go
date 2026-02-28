package repository

import "clean-architecture-learning/backend/internal/domain"

// TaskRepository はタスクの永続化に関する抽象。
type TaskRepository interface {
	Save(task domain.Task) error
	FindAll() ([]domain.Task, error)
	FindByID(id string) (domain.Task, error)
}

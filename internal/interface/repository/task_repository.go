package repository

import "clean-architecture-learning/internal/domain"

type TaskRepository interface {
	Save(task domain.Task) error
	FindAll() ([]domain.Task, error)
	FindByID(id string) (domain.Task, error)
}
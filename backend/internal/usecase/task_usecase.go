package usecase

import (
	"time"

	"github.com/google/uuid"

	"clean-architecture-learning/backend/internal/domain"
	"clean-architecture-learning/backend/internal/interface/repository"
)

// TaskUsecase はアプリ固有の操作手順（ユースケース）をまとめる。
type TaskUsecase struct {
	repo repository.TaskRepository
	now  func() time.Time
}

// NewTaskUsecase は依存を注入してユースケースを生成する。
func NewTaskUsecase(repo repository.TaskRepository, now func() time.Time) *TaskUsecase {
	return &TaskUsecase{repo: repo, now: now}
}

// Add はタスクを追加するユースケース。
func (u *TaskUsecase) Add(title string) (domain.Task, error) {
	id := newID()
	task, err := domain.NewTask(id, title, u.now())
	if err != nil {
		return domain.Task{}, err
	}
	if err := u.repo.Save(task); err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

// List はタスク一覧を取得するユースケース。
func (u *TaskUsecase) List() ([]domain.Task, error) {
	return u.repo.FindAll()
}

// Done はタスクを完了にするユースケース。
func (u *TaskUsecase) Done(id string) (domain.Task, error) {
	task, err := u.repo.FindByID(id)
	if err != nil {
		return domain.Task{}, err
	}
	task, err = task.Done(u.now())
	if err != nil {
		return domain.Task{}, err
	}
	if err := u.repo.Save(task); err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func newID() string {
	return uuid.New().String()
}

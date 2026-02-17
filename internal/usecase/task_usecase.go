package usecase

import (
	"time"

	"clean-architecture-learning/internal/domain"
	"clean-architecture-learning/internal/interface/repository"
)

type TaskUsecase struct {
	repo repository.TaskRepository
	now  func() time.Time
}

func NewTaskUsecase(repo repository.TaskRepository, now func() time.Time) *TaskUsecase {
	return &TaskUsecase{repo: repo, now: now}
}

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

func (u *TaskUsecase) List() ([]domain.Task, error) {
	return u.repo.FindAll()
}

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
	// temporary simple ID; will replace later
	return time.Now().Format("20060102150405.000000000")
}
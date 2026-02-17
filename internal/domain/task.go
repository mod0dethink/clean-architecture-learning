package domain

import "time"

// Status はドメイン上のタスク状態を表す。
type Status string

const (
	StatusOpen Status = "open"
	StatusDone Status = "done"
)

// Task は業務的な意味を持つドメインエンティティ。UI/DBに依存しない。
type Task struct {
	ID        string
	Title     string
	Status    Status
	CreatedAt time.Time
}

// NewTask はドメインの不変条件を満たすようにタスクを生成する。
func NewTask(id, title string, now time.Time) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyTitle
	}
	return Task{
		ID:        id,
		Title:     title,
		Status:    StatusOpen,
		CreatedAt: now,
	}, nil
}

// Done はドメインルールに従って完了状態にする。
func (t Task) Done(now time.Time) (Task, error) {
	if t.Status == StatusDone {
		return Task{}, ErrAlreadyDone
	}
	t.Status = StatusDone
	return t, nil
}

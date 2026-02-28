package domain

import "errors"

var (
	// ErrEmptyTitle はタイトルが空のときに返す。
	ErrEmptyTitle = errors.New("title is empty")
	// ErrAlreadyDone は既に完了しているタスクを完了にしようとしたときに返す。
	ErrAlreadyDone = errors.New("task already done")
)

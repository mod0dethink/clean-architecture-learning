package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"clean-architecture-learning/backend/internal/domain"
	"clean-architecture-learning/backend/internal/usecase"
)

type TaskHandler struct {
	uc *usecase.TaskUsecase
}

func NewTaskHandler(uc *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{uc: uc}
}

// GET /api/tasks
func (h *TaskHandler) ListTasks(c echo.Context) error {
	tasks, err := h.uc.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if tasks == nil {
		tasks = []domain.Task{}
	}
	return c.JSON(http.StatusOK, tasks)
}

// POST /api/tasks
type addTaskRequest struct {
	Title string `json:"title"`
}

func (h *TaskHandler) AddTask(c echo.Context) error {
	var req addTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}
	task, err := h.uc.Add(req.Title)
	if err != nil {
		if errors.Is(err, domain.ErrEmptyTitle) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, task)
}

// PUT /api/tasks/:id/done
func (h *TaskHandler) MarkDone(c echo.Context) error {
	id := c.Param("id")
	task, err := h.uc.Done(id)
	if err != nil {
		if errors.Is(err, domain.ErrAlreadyDone) {
			return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, task)
}

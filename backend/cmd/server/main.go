package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"

	"clean-architecture-learning/backend/internal/infrastructure/sqlite"
	"clean-architecture-learning/backend/internal/interface/handler"
	"clean-architecture-learning/backend/internal/usecase"
)

func main() {
	// --- Infrastructure ---
	db, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := sqlite.Migrate(db); err != nil {
		log.Fatal("migrate:", err)
	}

	// --- Dependency injection (innermost -> outermost) ---
	repo := sqlite.NewTaskRepository(db)
	uc := usecase.NewTaskUsecase(repo, time.Now)
	h := handler.NewTaskHandler(uc)

	// --- Echo setup ---
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// --- Routes ---
	api := e.Group("/api")
	api.GET("/tasks", h.ListTasks)
	api.POST("/tasks", h.AddTask)
	api.PUT("/tasks/:id/done", h.MarkDone)

	e.Logger.Fatal(e.Start(":8080"))
}

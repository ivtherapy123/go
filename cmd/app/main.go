package main

import (
	"awesomeProject1/internal/UserService"
	"awesomeProject1/internal/database"
	"awesomeProject1/internal/handlers"
	"awesomeProject1/internal/taskService"
	"awesomeProject1/internal/web/Users"
	"awesomeProject1/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal(err)
	}
	err2 := database.DB.AutoMigrate(&UserService.User{})
	if err2 != nil {
		log.Fatal(err)
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewService(taskRepo)
	taskHandler := handlers.NewHendler(taskService)

	userRepo := UserService.NewUserRepository(database.DB)
	userService := UserService.NewService(userRepo)
	userHandler := handlers.NewUserHendler(userService)
	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(taskHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	strictUserHandler := Users.NewStrictHandler(userHandler, nil)
	Users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

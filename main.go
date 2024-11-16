package main

import (
	"log"
	"todoList/config"
	"todoList/pkg/app"
	"todoList/pkg/http"
	"todoList/pkg/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// DB 초기화
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	// 애플리케이션 초기화
	app := app.NewApplication(db)

	// 라우터 초기화 및 설정
	router := http.NewRouter(app)
	router.SetupRoutes()
	// 서버 시작
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"todoList/pkg/app"
)

type Router struct {
	engine *gin.Engine
	app    *app.Application
}

func NewRouter(app *app.Application) *Router {
	engine := gin.Default()

	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	return &Router{
		engine: engine,
		app:    app,
	}
}

func (r *Router) SetupRoutes() {
	api := r.engine.Group("/api")
	v1 := api.Group("/v1")

	NewUserRouter(v1, r.app.UserModule).Setup()
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}

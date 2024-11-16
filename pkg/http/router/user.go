package router

import (
	"github.com/gin-gonic/gin"
	"todoList/pkg/user"
)

type UserRouter struct {
	group      *gin.RouterGroup
	controller *user.Controller
}

func NewUserRouter(group *gin.RouterGroup, module *user.Module) *UserRouter {
	return &UserRouter{
		group:      group,
		controller: module.Controller(),
	}
}

func (r *UserRouter) Setup() {
	users := r.group.Group("/users")
	{

		users.GET("/:id", r.controller.Get)

	}
}

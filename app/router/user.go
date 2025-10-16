package router

import (
	"xm/internal/entity/request"
	"xm/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (c *Capsule) UserRoutes(r *gin.RouterGroup) {
	userHandler := c.Handler.UserHandler

	userV1 := r.Group("/user")

	userV1.POST(
		"/register",
		middleware.Validate[request.UserRequest]("CreateUserRequest"),
		userHandler.HandleCreateUser)

	userV1.POST(
		"/login",
		middleware.Validate[request.UserRequest]("LoginUserRequest"),
		userHandler.HandleLoginUser)
}

package middleware

import (
	"net/http"
	"xm/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Validate[T any](tag string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestModel T
		err := c.BindJSON(&requestModel)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, utils.Fail(100, err.Error()))
			c.Abort()
		}
		validate := validator.New()
		if err := validate.Struct(&requestModel); err != nil {
			c.IndentedJSON(http.StatusBadRequest, utils.Fail(100, err.Error()))
			c.Abort()
		}
		c.Set(tag, requestModel)
	}
}

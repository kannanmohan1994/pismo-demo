package handler

import (
	"context"
	"net/http"
	"pismo/utils"

	"github.com/gin-gonic/gin"
)

// Authorization.
type RequestHandler interface {
	Decode(c *gin.Context) error
}

func handleRequest(c *gin.Context, req RequestHandler, service func(context.Context, RequestHandler) (httpCode int, data interface{}, err error)) {
	var err error
	var statusCode int
	var data interface{}

	defer func() {
		if err != nil {
			c.Errors = append(c.Errors, &gin.Error{Err: err})
		}
		c.JSON(statusCode, utils.Send(data, err, ""))
	}()

	if err = req.Decode(c); err != nil {
		statusCode = http.StatusBadRequest
		return
	}
	statusCode, data, err = service(c, req)
	if err != nil {
		if statusCode == 0 {
			statusCode = http.StatusBadRequest
		}
	} else {
		if statusCode == 0 {
			statusCode = http.StatusOK
		}
	}
}

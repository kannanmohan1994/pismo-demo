package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"pismo/consts"
	"pismo/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (m *middleware) Trace(skippedURLs ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := fmt.Sprintf("%s: [%s]", c.Request.Method, c.Request.URL.Path)
		start := time.Now().UTC()
		for _, url := range skippedURLs {
			if c.Request.URL.Path == url {
				return
			}
		}

		requestBody, err := c.GetRawData()
		if err != nil {
			m.logger.Errorf("failed to get reqeuest body")
			return
		} else {
			c.Request.Body = io.NopCloser(io.Reader(bytes.NewBuffer(requestBody)))
		}

		var ctx context.Context = c

		setAuthHeaders(c)
		err = setCorrelationID(c)
		if err != nil {
			m.logger.Errorf("failed to set correlation-id")
			return
		}

		log := m.logger.WithContext(ctx)
		var data = make([]interface{}, 0)
		data = append(data, "request", string(requestBody))
		log.Infow(endpoint+" request started", data...)
		defer func() {
			data = append(data, "ts", time.Since(start))
			var endPointErr error
			if len(c.Errors) != 0 {
				endPointErr = c.Errors[0].Err
				log.Errorf(fmt.Sprintf("%s request finished with error: %s", endpoint, endPointErr.Error()), data...)
			} else {
				log.Infow(endpoint+" request finished", data...)
			}
		}()

		c.Next()
	}
}

func setCorrelationID(c *gin.Context) error {
	correlationID, ok := c.Get(consts.CorrelationID)
	if !ok || !utils.VerifyUUID(correlationID.(string)) {
		uuid, err := utils.GenerateUUID()
		if err != nil {
			return errors.Wrap(err, "failed to generate new correlation-id")
		}
		c.Set(consts.CorrelationID, uuid)
	}
	return nil
}

func setAuthHeaders(c *gin.Context) {
	c.Set("Authorization", (c.Request.Header.Get("Authorization")))
}

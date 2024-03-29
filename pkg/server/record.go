package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Record() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		start := time.Now()

		defer func() {
			cost := time.Since(start).Microseconds()
			httpCode := ctx.Writer.Status()
			clientIP := ctx.ClientIP()
			clientUA := ctx.Request.UserAgent()
			method := ctx.Request.Method

			entry := logrus.WithFields(logrus.Fields{
				"cost":   cost,
				"method": method,
				"code":   httpCode,
				"ip":     clientIP,
				"path":   path,
			})

			if len(ctx.Errors) > 0 {
				entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				msg := fmt.Sprintf("user-agent: %s", clientUA)
				switch {
				case httpCode >= http.StatusInternalServerError:
					entry.Error(msg)
				case httpCode >= http.StatusBadRequest:
					entry.Warn(msg)
				default:
					entry.Info(msg)
				}
			}
		}()

		ctx.Next()
	}
}

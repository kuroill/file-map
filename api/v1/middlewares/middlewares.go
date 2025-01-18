package middlewares

import (
	"file-map-server/api/v1/model"
	"file-map-server/log"
	"file-map-server/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	*model.Model
}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Cache-Control")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api") {
			switch path {
			case "/api/file/auth-pwd":
				c.Next()
				return
			case "/api/file/stream-video":
				c.Next()
				return
			case "/api/file/download":
				c.Next()
				return
			default:

				err := utils.AuthToken(c.GetHeader("Authorization"))
				if err != nil {
					m.ErrorRes(c, http.StatusUnauthorized, err.Error())
					c.Abort()
					return
				}
				c.Next()
			}
		} else {
			c.Next()
			return
		}
	}
}

func (m *Middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiPath := c.Request.URL.Path
		queryParams := c.Request.URL.Query()
		params := c.Params

		log.Info("Request: ", c.Request.Method, apiPath)
		log.Info("Query Params: ", queryParams)
		log.Info("Route Params: ", params)

		c.Next()

		statusCode := c.Writer.Status()
		log.Info("Response status: ", statusCode)
	}
}

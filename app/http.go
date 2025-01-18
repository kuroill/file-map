package app

import (
	"file-map-server/api/v1/middlewares"
	"file-map-server/api/v1/router/file"
	"file-map-server/app/config"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	m := middlewares.New()
	r.Use(m.Cors())
	r.Use(m.AuthMiddleware())
	r.Use(m.Logger())

	path := "../app/dist"
	if os.Getenv("GIN_MODE") == "release" {
		path = "/app/dist"
	}

	r.Static("/assets", filepath.Join(path, "assets"))
	r.LoadHTMLFiles(filepath.Join(path, "index.html"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	file.New(r)

	port := config.GetConfig().Server.HTTP.Port
	r.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

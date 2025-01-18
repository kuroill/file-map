package file

import (
	"file-map-server/utils"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StreamVideo(c *gin.Context) {
	path := c.Query("path")
	token := c.Query("token")

	err := utils.AuthToken(token)
	if err != nil {
		h.ErrorRes(c, http.StatusUnauthorized, err.Error())
		return
	}

	file, err := os.Open(path)
	if err != nil {
		h.ErrorRes(c, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		h.ErrorRes(c, http.StatusInternalServerError, err.Error())
		return
	}

	mimeType := "application/octet-stream"
	if ext := filepath.Ext(path); ext != "" {
		mimeType = mime.TypeByExtension(ext)
	}

	c.Header("Content-Type", mimeType)
	c.Header("Content-Disposition", "inline; filename="+filepath.Base(path))
	http.ServeContent(c.Writer, c.Request, filepath.Base(path), fileInfo.ModTime(), file)
}

package file

import (
	"file-map-server/utils"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Download(c *gin.Context) {
	path := c.Query("path")
	token := c.Query("token")

	err := utils.AuthToken(token)
	if err != nil {
		h.ErrorRes(c, http.StatusUnauthorized, err.Error())
		return
	}

	file, err := os.Open(path)
	if err != nil {
		h.ErrorRes(c, http.StatusBadRequest, "file not found")
		return
	}
	defer file.Close()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename*=UTF-8''"+url.QueryEscape(filepath.Base(path)))
	c.Header("Content-Type", "application/octet-stream")
	c.File(path)
}

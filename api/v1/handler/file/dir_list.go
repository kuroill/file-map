package file

import (
	"file-map-server/api/v1/model"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Dirlist(c *gin.Context) {
	path := c.Query("path")

	dl, err := listFiles(path)
	if err != nil {
		h.ErrorRes(c, http.StatusBadRequest, err.Error())
		return
	}
	h.SuccessRes(c, http.StatusOK, dl)
}

func listFiles(dir string) (*[]model.DirList, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dl []model.DirList
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		dl = append(dl, model.DirList{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			IsDir:   entry.IsDir(),
			CurDir:  dir,
		})
	}

	return &dl, nil
}

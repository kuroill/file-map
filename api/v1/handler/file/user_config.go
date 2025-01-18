package file

import (
	"file-map-server/api/v1/model"
	"file-map-server/app/config"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserConfig(c *gin.Context) {
	rd := config.GetConfig().User.RootDir
	goos := ""

	mode := os.Getenv("GIN_MODE") == "release"
	if mode {
		rd = config.GetConfig().User.Mnt
	} else {
		goos = runtime.GOOS
	}

	h.SuccessRes(c, http.StatusOK, model.UserConfig{
		RootDir: rd,
		Os:      goos,
		Mode:    mode,
	})
}

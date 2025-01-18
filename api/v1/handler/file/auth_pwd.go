package file

import (
	"file-map-server/app/config"
	"file-map-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthPwd(c *gin.Context) {
	password := c.Query("password")

	if len(password) <= 0 || config.GetConfig().User.Password != password {
		h.ErrorRes(c, http.StatusBadRequest, "password error")
		return
	}

	secret := config.GetConfig().User.SecretKey
	encrypt, err := utils.EncryptAES([]byte(secret), []byte(password))
	if err != nil {
		h.ErrorRes(c, http.StatusUnauthorized, err.Error())
	}

	h.SuccessRes(c, http.StatusOK, encrypt)
}

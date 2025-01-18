package model

import (
	"file-map-server/log"
	"time"

	"github.com/gin-gonic/gin"
)

type Model struct{}

type UserConfig struct {
	RootDir string `json:"root_dir"`
	Os      string `json:"os"`
	Mode    bool   `json:"mode"`
}

type DirList struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"mod_time"`
	IsDir   bool      `json:"is_dir"`
	CurDir  string    `json:"cur_dir"`
}

type ResModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Res(code int, message string, data interface{}) *ResModel {
	return &ResModel{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (h *Model) ErrorRes(c *gin.Context, code int, message string) {
	log.Error(message)
	c.JSON(code, Res(code, message, nil))
}

func (h *Model) SuccessRes(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Res(code, "", data))
}

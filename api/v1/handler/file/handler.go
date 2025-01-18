package file

import "file-map-server/api/v1/model"

type Handler struct {
	*model.Model
}

func New() *Handler {
	return &Handler{}
}

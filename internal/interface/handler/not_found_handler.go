package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundHandler interface {
	NotFound(c *gin.Context)
}

type notFoundHandler struct {
}

func NewNotFoundHandler() NotFoundHandler {
	return &notFoundHandler{}
}

func (h *notFoundHandler) NotFound(c *gin.Context) {
	response := map[string]interface{}{
		"result": nil,
		"metadata": map[string]string{
			"error": "Resouce not found",
		},
	}
	c.JSON(http.StatusNotFound, gin.H(response))
}

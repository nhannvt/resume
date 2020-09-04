package handler

import (
	"net/http"

	"github.com/nhannvt/resume/internal/usecase"
	"github.com/gin-gonic/gin"
)

const (
	statusHealthy  = "healty"
	statusUnhealty = "unhealty"
)

type HealthCheckHandler interface {
	HealthCheck(c *gin.Context)
}

// healthCheckHandler has a function to check own health status.
type healthCheckHandler struct {
	u usecase.HealthCheckUseCase
}

// NewHealthCheckHandler creates new healthCheckHandler struct.
func NewHealthCheckHandler(u usecase.HealthCheckUseCase) HealthCheckHandler {
	return &healthCheckHandler{u}
}

// HealthCheck checks if sforum is available and writes result to response body via gin context.
func (h *healthCheckHandler) HealthCheck(c *gin.Context) {

	err := h.u.HealthCheck()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": statusUnhealty, "metadata": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": statusHealthy, "metadata": nil})
}

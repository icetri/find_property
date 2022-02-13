package v1

import (
	"github.com/find_property/internal/domain"
	"github.com/find_property/internal/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// @Summary Init
// @Description ПИНГ
// @ID ping
// @Success 200 {integer} integer 1
// @Failure 400 {integer} integer 0
// @Router /ping [get]
func (h *Handler) Init(c *gin.Context) {

	c.String(http.StatusOK, "pong")

}

type result struct {
	Err string `json:"error"`
}

func apiErrorEncode(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	var code int
	if customError, ok := err.(*domain.CustomError); ok {
		code = customError.Code
		c.Writer.WriteHeader(customError.Code)
	}

	r := result{Err: err.Error()}

	c.AbortWithStatusJSON(code, r)
}

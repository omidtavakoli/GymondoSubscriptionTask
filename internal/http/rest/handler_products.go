package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Products(c *gin.Context) {
	res, err := h.SubscriptionService.GetProductsList()
	if err != nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
package rest

import (
	"Gymondo/internal/subscription"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func (h *Handler) DummyDataGenerator(c *gin.Context) {
	h.SubscriptionService.UserGenerator(10)
	h.SubscriptionService.ProductGenerator(10)
	h.SubscriptionService.PlanGenerator()

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) Products(c *gin.Context) {
	res, err := h.SubscriptionService.GetProductsList()
	if err != nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Product(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, nil)
		return
	}
	res, err := h.SubscriptionService.GetProductById(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, res)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) BuyProduct(c *gin.Context) {
	var buyRequest subscription.BuyRequest
	if err := c.ShouldBind(&buyRequest); err != nil {
		if vErr, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, GetFailResponseFromValidationErrors(vErr))
		}
		return
	}
	plan, bpErr := h.SubscriptionService.BuyProduct(buyRequest)
	if bpErr != nil {
		c.JSON(http.StatusNotFound, plan)
		return
	}
	c.JSON(http.StatusOK, plan)
}

func (h *Handler) FetchPlansByUserId(c *gin.Context) {
	id := c.Param("user_id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	plan, bpErr := h.SubscriptionService.FetchPlansByUserId(userId)
	if bpErr != nil {
		c.JSON(http.StatusNotFound, plan)
		return
	}
	c.JSON(http.StatusOK, plan)
}

func (h *Handler) ChangePlanStatus(c *gin.Context) {
	var changeStatus subscription.ChangeStatus
	if err := c.ShouldBind(&changeStatus); err != nil {
		if vErr, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, GetFailResponseFromValidationErrors(vErr))
		}
		return
	}
	//todo: make better response body
	bpErr := h.SubscriptionService.ChangeUserPlanStatus(changeStatus)
	if bpErr != nil {
		c.JSON(http.StatusNotFound, "done")
		return
	}
	c.JSON(http.StatusOK, "done")
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"github.com/labstack/gommon/log"
)

type PurchaseController struct {
	logger *log.Logger
}

func NewPurchaseController(logger *log.Logger) *PurchaseController {
	return &PurchaseController {
		logger: logger,
	}
}

func (p *PurchaseController) Post(c *gin.Context, insert func(entity.Purchase) (entity.Purchase, error)) {
	var purchase entity.Purchase
	if err := c.BindJSON(&purchase); err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, c.Error(err))
		return
	}

	res, err := insert(purchase)
	if err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}
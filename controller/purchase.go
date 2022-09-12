package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"github.com/labstack/gommon/log"
)

type PurchaseController struct {
	logger *log.Logger
}

func NewPurchaseController(logger *log.Logger) *PurchaseController {
	return &PurchaseController{
		logger: logger,
	}
}

func (p *PurchaseController) Post(c *gin.Context, insert func(*entity.Purchase) error) {
	var purchase entity.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, c.Error(err))
		return
	}

	err := insert(&purchase)
	if err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusCreated, purchase)
}

func (p *PurchaseController) PutToggle(c *gin.Context, find func(int) (entity.Purchase, error)) {
	var id int
	if err := c.Bind(&id); err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	res, err := find(id)
	if err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (p *PurchaseController) GetProductsOne(c *gin.Context, find func(int) (entity.Purchase, error)) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, c.Error(err))
		return
	}

	res, err := find(id)
	if err != nil {
		p.logger.Error(err)
		c.Copy().AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

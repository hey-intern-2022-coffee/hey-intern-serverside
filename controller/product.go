package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"github.com/labstack/gommon/log"
)

type ProductController struct {
	logger *log.Logger
}

func NewProductController(logger *log.Logger) *ProductController {
	return &ProductController{
		logger: logger,
	}
}

func (p *ProductController) Post(c *gin.Context, insert func(*entity.Product) error) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, c.Error(err))
		return
	}

	if err := insert(&product); err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) GetOne(c *gin.Context, find func(int)(*entity.Product, error)) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
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

func (p *ProductController) GetAll(c *gin.Context, find func() ([]entity.Product, error)) {
	res, err := find()
	if err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

func (p *ProductController) PatchPurchase(c *gin.Context, patch func(int) (*entity.Product, error)) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	res, err := patch(product.ID)
	if err != nil {
		p.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, c.Error(err))
		return
	}

	c.JSON(http.StatusOK, res)
}

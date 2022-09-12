package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/config"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/controller"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/infra"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/log"
)

func main() {
	_, err := infra.NewDB()
	if err != nil {
		os.Exit(1)
	}

	log := log.New()
	purchaseCtrl := controller.NewPurchaseController(log)
	productCtrl := controller.NewProductController(log)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			config.Arrow(),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
			"PATCH",
		},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.POST("/purchase", func(ctx *gin.Context) {
		purchaseCtrl.Post(ctx, func(p *entity.Purchase) error {
			return nil
		})
	})

	r.PATCH("/purchase", func(c *gin.Context) {
		productCtrl.PatchPurchase(c, func(i int) (entity.Product, error) {
			return entity.Product{}, nil
		})
	})

	r.PUT("/purchase", func(ctx *gin.Context) {
		purchaseCtrl.PutToggle(ctx, func(i int) (entity.Purchase, error) {
			return entity.Purchase{}, nil
		})

		r.GET("/purchase/:id", func(ctx *gin.Context) {
			purchaseCtrl.GetProductsOne(ctx, func(i int) (entity.Purchase, error) {
				return entity.Purchase{}, nil
			})
		})

		r.POST("/product", func(ctx *gin.Context) {
			productCtrl.Post(ctx, func(p *entity.Product) error {
				return nil
			})
		})

		r.GET("/products", func(ctx *gin.Context) {
			productCtrl.GetAll(ctx, func() ([]entity.Product, error) {
				return []entity.Product{}, nil
			})
		})

		r.GET("/onlinestore/allproducts", func(ctx *gin.Context) {
			productCtrl.GetAll(ctx, func() ([]entity.Product, error) {
				return []entity.Product{}, nil
			})
		})
	})

	r.Run(":8080")
}

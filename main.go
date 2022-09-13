package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/config"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/controller"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/infra"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/log"
)

func main() {
	db, err := infra.NewDB()
	if err != nil {
		os.Exit(1)
	}

	productRepo := infra.NewProductRepository(db)
	purchaseRepo := infra.NewPurchaseRepository(db)

	log := log.New()
	purchaseCtrl := controller.NewPurchaseController(log)
	productCtrl := controller.NewProductController(log)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			config.Arrow(),
			"https://stores-coffee-manager.netlify.app",
			"https://stores-coffee.netlify.app",
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
		purchaseCtrl.Post(ctx, purchaseRepo.Insert)
	})

	r.PATCH("/purchase", func(c *gin.Context) {
		productCtrl.PatchPurchase(c, productRepo.PatchPurchase)
	})

	r.PATCH("/purchase/delivered", func(ctx *gin.Context) {
		purchaseCtrl.PutToggle(ctx, purchaseRepo.ToggleIsAcceptance)
	})

	r.GET("/purchase/:id", func(ctx *gin.Context) {
		purchaseCtrl.GetPurchaseOne(ctx, purchaseRepo.FindByPurchaseID)
	})

	r.POST("/product", func(ctx *gin.Context) {
		productCtrl.Post(ctx, productRepo.Insert)
	})

	r.GET("/products", func(ctx *gin.Context) {
		productCtrl.GetAll(ctx, productRepo.FindAll)
	})

	r.GET("/onlinestore/allproducts", func(ctx *gin.Context) {
		productCtrl.GetAll(ctx, productRepo.FindAll)
	})

	r.GET("/products/:id", func(ctx *gin.Context) {
		productCtrl.GetOne(ctx, productRepo.FindIdOne)
	})

	r.Run(config.Port())
}

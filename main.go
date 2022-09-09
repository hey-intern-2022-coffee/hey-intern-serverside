package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/infra"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/controller"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/config"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/domain/entity"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/log"
)

func main() {
	_, err := infra.NewDB()
	if err != nil {
		os.Exit(1)
	}

	log := log.New()
	purchaseCtrl := controller.NewPurchaseController(log)

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
		},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.POST("/purchases", func(ctx *gin.Context) {
		purchaseCtrl.Post(ctx, func(p entity.Purchase) (entity.Purchase, error) {
			return entity.Purchase{}, nil
		})
	})

	r.Run(":8080")
}

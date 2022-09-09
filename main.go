package main

import (
	"os"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/config"
	"github.com/hey-intern-2022-coffee/hey-intern-serverside/infra"
)

func main() {

	_, err := infra.NewDB()
	if err != nil {
		os.Exit(1)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
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

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "healthcheck success") })
	r.Run(config.Port())
}

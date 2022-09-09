package main

import(
	"net/http"

	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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
}
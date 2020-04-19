package main

import (
	"time"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.New()
	p := ginprometheus.NewPrometheus("go_sli_slo_app")
	p.Use(r)

	r.GET("/fast_response", func(c *gin.Context) {
		c.JSON(200, "Fast Response :)")
	})

	r.GET("/slow_response", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(200, "Slow Response :(")
	})

	r.GET("/error_response", func(c *gin.Context) {
		c.JSON(500, "Error Response")
	})

	r.Run(":2112")
}

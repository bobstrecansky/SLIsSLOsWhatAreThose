package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.New()
	// Optional custom metrics list
	customMetrics := []*ginprometheus.Metric{
		&ginprometheus.Metric{
			ID:          "1234",          // optional string
			Name:        "histogram_vec", // required string
			Description: "histogram_vec", // required string
			Type:        "histogram_vec", // required string
		},
	}
	p := ginprometheus.NewPrometheus("go_sli_slo_app", customMetrics)

	p.Use(r)

	r.GET("/fast_response", func(c *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		time.Sleep(time.Duration(time.Duration(n) * time.Millisecond))
		c.JSON(200, gin.H{"Response Time (ms)": n})
	})

	r.GET("/slow_response", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(200, "Slow Response")
	})

	r.GET("/error_response", func(c *gin.Context) {
		c.JSON(500, "Error Response")
	})

	r.Run(":2112")
}

package main

import (
	"synthtic-harbor/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/v0/dental/inquiries/:id", handlers.GetDentalInquiry)
	r.GET("/api/v0/dental/inquiries", handlers.GetDentalInquiries)

	r.Run(":8000") // Listen and serve on 0.0.0.0:8000
}

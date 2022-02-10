package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize engine
	r := gin.Default()
	// Initialize database
	database := NewDatabase()
	// Initialize routes
	InitRoutes(r, database)
	// Run engine
	r.Run()
}

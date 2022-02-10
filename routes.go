package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Initialize Routes
func InitRoutes(r *gin.Engine, d *Database) {
	// Access URL route
	r.GET("/:token", func(c *gin.Context) {
		url, ok := d.Get(c.Param("token"))
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid short url"})
			return
		}
		c.Redirect(http.StatusMovedPermanently, url)
	})

	// Ranking route
	r.GET("/ranking", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"top_100": d.GetTop100()})
	})

	// Register URL route
	r.POST("/", func(c *gin.Context) {
		url := c.Request.URL.Query().Get("url")
		re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
		if !re.MatchString(url) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
			return
		}

		id := d.Length()
		token := NewRandomToken(id)
		d.Push(url, token)

		c.JSON(http.StatusOK, gin.H{"short": "http://localhost:8080/" + token, "long": url, "id": id})
	})
}

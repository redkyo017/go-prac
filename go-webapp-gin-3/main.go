package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var f embed.FS

func main() {
	router := gin.Default()
	// router.StaticFile("/", "./public/index.html")

	// router.Static("/public", "./public")

	// router.StaticFS("/fs", http.FileSystem(http.FS(f)))
	// router.GET("/employee", func(c *gin.Context) {
	// 	c.File("./public/employee.html")
	// })

	// router.POST("/employee", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "New request POSTed successfully!")
	// })
	// router.GET("/employee/:username/*rest", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"username": c.Param("username"),
	// 		"rest":     c.Param("rest"),
	// 	})
	// })
	adminGroup := router.Group("/admin")
	adminGroup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "administrator users")
	})

	adminGroup.GET("/roles", func(c *gin.Context) {
		c.String(http.StatusOK, "administrator roles")
	})

	adminGroup.GET("/policies", func(c *gin.Context) {
		c.String(http.StatusOK, "administrator policies")
	})

	log.Fatal(router.Run(":3000"))
}

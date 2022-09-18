package main

import (
	"log"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func v1EndpointHandler(c *gin.Context) {
	c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
}

func v2EndpointHandler(c *gin.Context) {
	c.String(200, "v2: %s %s", c.Request.Method, c.Request.URL.Path)
}

func add(c *gin.Context) {
	var ap AddParams
	if err := c.ShouldBindJSON(&ap); err != nil {
		c.JSON(400, gin.H{"error": "Calculation error"})
		return
	}
	c.JSON(200, gin.H{"answer": ap.X + ap.Y})
}

type AddParams struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Product struct {
	Id   int    `json:"id" xml:"Id" yaml:"id"`
	Name string `json:"name" xml:"Name" yaml:"name"`
}

func FindUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.GetHeader("User-Agent"))
		// before
		c.Next()
		// After
	}
}

func main() {
	// router = gin.Default()
	// // router.Use(cors.Default())
	// router.Use(FindUserAgent())
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "CORS works"})
	// })
	// router.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello world",
	// 	})
	// })
	// router.GET("/os", func(c *gin.Context) {
	// 	c.JSON(200, runtime.GOOS)
	// })

	// router.POST("/add", add)
	// router.GET("/productJSON", func(c *gin.Context) {
	// 	product := Product{1, "Apple"}
	// 	c.JSON(200, product)
	// })
	// router.GET("/productXML", func(c *gin.Context) {
	// 	product := Product{2, "Banana"}
	// 	c.XML(200, product)
	// })
	// router.GET("/productYAML", func(c *gin.Context) {
	// 	product := Product{3, "Mango"}
	// 	c.YAML(200, product)
	// })

	// v1 := router.Group("/v1")
	// v1.GET("/products", v1EndpointHandler)
	// v1.GET("/products/:productId", v1EndpointHandler)
	// v1.POST("/products", v1EndpointHandler)
	// v1.PUT("/products/:productId", v1EndpointHandler)
	// v1.DELETE("/products/:productId", v1EndpointHandler)

	// v2 := router.Group("/v2")

	// v2.GET("/products", v2EndpointHandler)
	// v2.GET("/products/:productId", v2EndpointHandler)
	// v2.POST("/products", v2EndpointHandler)
	// v2.PUT("/products/:productId", v1EndpointHandler)
	// v2.DELETE("/products/:productId", v1EndpointHandler)

	// router.Run(":5001")
	// print_service()

	invoice_generate()
}

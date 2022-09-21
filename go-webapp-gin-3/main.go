package main

import (
	"embed"
	"log"
	"net/http"
	"strconv"
	"time"

	// "github.com/gin-gonic/binding"
	"github.com/gin-gonic/gin"
)

var f embed.FS

type TimeoffRequest struct {
	Date   time.Time `json:"date" form:"date" binding:"required" time_format:"2006-01-02"`
	Amount float64   `json:"amount" form:"amount" binding:"required,gt=0"`
}

// var ValidatorFuture validator.Func = func(fl validator.FieldLevel) bool {
// 	date, ok := fl.Field().Interface().(time.Time)
// 	if ok {
// 		return date.After(time.Now())
// 	}
// 	return true
// }

func main() {
	router := gin.Default()

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("future", ValidatorFuture)
	// }
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
	// adminGroup := router.Group("/admin")
	// adminGroup.GET("/users", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "administrator users")
	// })

	// adminGroup.GET("/roles", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "administrator roles")
	// })

	// adminGroup.GET("/policies", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "administrator policies")
	// })
	// router.GET("/*rest", func(c *gin.Context) {
	// 	url := c.Request.URL.String()
	// 	headers := c.Request.Header
	// 	cookies := c.Request.Cookies()

	// 	c.IndentedJSON(http.StatusOK, gin.H{
	// 		"url":     url,
	// 		"headers": headers,
	// 		"cookies": cookies,
	// 	})
	// })
	router.GET("/query/*rest", func(c *gin.Context) {
		username := c.Query("username")
		year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := c.QueryArray("month")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})
	})

	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})
	router.POST("/employee", func(c *gin.Context) {
		// date := c.PostForm("date")
		// amount := c.PostForm("amount")
		// username := c.DefaultPostForm("username", "me")
		// c.IndentedJSON(http.StatusOK, gin.H{
		// 	"date":     date,
		// 	"amount":   amount,
		// 	"username": username,
		// })
		var timeoffRequest TimeoffRequest
		if err := c.ShouldBind(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}

	})

	apiGroup := router.Group("/api")
	apiGroup.POST("/timeoff", func(c *gin.Context) {
		var timeoffRequest TimeoffRequest
		if err := c.ShouldBindJSON(&timeoffRequest); err == nil {
			c.JSON(http.StatusOK, timeoffRequest)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}

	})

	log.Fatal(router.Run(":3000"))
}

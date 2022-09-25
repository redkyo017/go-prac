package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-webapp-gin-3/employee"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Use(gin.BasicAuth(gin.Accounts{"admin": "password"}))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(logErrorMiddleware)
	r.Use(gin.CustomRecovery(myrecoveryFunc))
	registerRoutes(r)

	r.Run()

}

func registerRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		// c.File("./public/index.html")
		c.Redirect(http.StatusTemporaryRedirect, "/employees")
	})

	r.GET("/employees", func(c *gin.Context) {
		// c.File("./public/employee.html")
		c.HTML(http.StatusOK, "index.tmpl", employee.GetAll())
	})

	r.GET("/employees/:employeeID", func(c *gin.Context) {
		employeeIDRaw := c.Param("employeeID")
		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			log.Println(*emp)
			c.HTML(http.StatusOK, "employee.tmpl", *emp)
		}
	})

	r.POST("/employees/:employeeID", func(c *gin.Context) {
		var timeoff employee.TimeOff
		err := c.ShouldBind(&timeoff)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		timeoff.Type = employee.TimeoffTypePTO
		timeoff.Status = employee.TimeoffStatusRequested

		employeeIDRaw := c.Param("employeeID")
		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			emp.TimeOff = append(emp.TimeOff, timeoff)
			c.Redirect(http.StatusFound, "/employees/"+employeeIDRaw)
		}
	})
	g := r.Group("/api/employees", Benmark)
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, employee.GetAll())
	})
	g.GET("/:employeeID", func(c *gin.Context) {
		employeeIDRaw := c.Param("employeeID")
		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			c.JSON(http.StatusOK, *emp)
		}
	})
	g.POST("/:employeeID", func(c *gin.Context) {
		var timeoff employee.TimeOff
		err := c.ShouldBindJSON(&timeoff)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		timeoff.Type = employee.TimeoffTypePTO
		timeoff.Status = employee.TimeoffStatusRequested

		employeeIDRaw := c.Param("employeeID")
		if emp, ok := tryToGetEmployee(c, employeeIDRaw); ok {
			emp.TimeOff = append(emp.TimeOff, timeoff)
			c.JSON(http.StatusOK, *emp)
		}
	})

	r.GET("/errors", func(c *gin.Context) {
		err := &gin.Error{
			Err:  errors.New("something went horribly wrong"),
			Type: gin.ErrorTypeRender | gin.ErrorTypePublic,
			Meta: "this error was intentional",
		}
		c.Error(err)
	})

	r.GET("panic", func(c *gin.Context) {
		panic("a Go program should almostnever call 'panic'")
	})

	r.Static("/public", "./public")
}

func tryToGetEmployee(c *gin.Context, IDRaw string) (*employee.Employee, bool) {
	employeeId, err := strconv.Atoi(IDRaw)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return nil, false
	}
	emp, err := employee.Get(employeeId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil, false
	}
	return emp, true
}

var Benmark gin.HandlerFunc = func(c *gin.Context) {
	t := time.Now()

	c.Next()

	elapsed := time.Since(t)
	log.Println("Time to process", elapsed)
}

var logErrorMiddleware = func(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		log.Println(map[string]any{
			"err":  err.Error(),
			"type": err.Type,
			"meta": err.Meta,
		})
	}
}

var myrecoveryFunc gin.RecoveryFunc = func(c *gin.Context, err any) {
	log.Println("Custom recovery functions can be  used to add fine-grained control over recovery strategies")
}

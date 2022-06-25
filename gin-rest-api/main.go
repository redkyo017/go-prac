package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func main() {
	fmt.Println("con co be be")
	log.Println("con meo bay bay")
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
<<<<<<< HEAD
=======

>>>>>>> 81ce7fc2ce05890038cabe81c724440ba9ba77f3
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.POST("/books", func(c *gin.Context) {
		var book Book
<<<<<<< HEAD
		if err := c.ShouldBindJSON(&book); err != nil {
			log.Println("hehee", err)
=======

		if err := c.ShouldBindJSON(&book); err != nil {
>>>>>>> 81ce7fc2ce05890038cabe81c724440ba9ba77f3
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, book)
		c.JSON(http.StatusCreated, books)
	})
<<<<<<< HEAD
=======

	r.DELETE("/books/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, b := range books {
			if b.ID == id {
				books = append(books[:i], books[i+1:]...)
				break
			}
		}
		c.JSON(http.StatusOK, books)
	})
>>>>>>> 81ce7fc2ce05890038cabe81c724440ba9ba77f3
	r.Run()
}

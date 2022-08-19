package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	c.HTML(
		// set HTTP status
		http.StatusOK,
		"index.html",
		// pass data
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

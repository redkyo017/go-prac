package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type Invoice struct {
	InvoiceId   int    `json:"invoiceId"`
	CustomerId  int    `json:"customerId" binding:"required,gte=0"`
	Price       int    `json:"price" binding:"required,gte=0"`
	Description string `json:"description" binding:"required"`
}

type PrintJob struct {
	Format    string `json:"format" binding:"required"`
	InvoiceId int    `json:"invoiceId" binding:"required,gte=0"`
	JobId     int    `json:"jobId" binding:"gte=0"`
}

func createPrintJob(invoiceId int) {
	client := resty.New()
	var p PrintJob
	// Call PrinterService via RESTful interface
	_, err := client.R().
		SetBody(PrintJob{Format: "A4", InvoiceId: invoiceId}).
		SetResult(&p).
		Post("http://localhost:5002/print-jobs")
	log.Println("con heo", p)
	if err != nil {
		log.Println("InvoiceGenerator: unable to connect PrinterService")
		return
	}
	log.Printf("InvoiceGenerator: created print job #%v via PrinterService", p.JobId)
}

func invoice_generate() {
	router := gin.Default()
	router.POST("/invoices", func(c *gin.Context) {
		var iv Invoice
		if err := c.ShouldBindJSON(&iv); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!"})
			return
		}
		log.Println("InvoiceGenerator: creating new invoice...")
		rand.Seed(time.Now().UnixNano())
		iv.InvoiceId = rand.Intn(1000)
		log.Printf("InvoiceGenerator: created invoice #%v", iv.InvoiceId)

		createPrintJob(iv.InvoiceId) // Ask PrinterService to create a print job
		c.JSON(200, iv)
	})
	router.Run(":5003")
}

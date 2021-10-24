package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var thisDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func main() {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("finance.db"), &gorm.Config{})
	check(err)

	// Migrate tables
	err = db.AutoMigrate(
		&Payment{},
		&Category{},
	)
	check(err)

	// Load templates
	router := gin.Default()
	router.LoadHTMLGlob(path.Join(thisDir, "template/*.html"))

	// Serve static files or 404
	router.NoRoute(func(c *gin.Context) {
		resource := c.Request.URL.Path
		staticFile := path.Join(thisDir, "static", resource)
		if info, err := os.Stat(staticFile); err == nil && !info.IsDir() {
			c.File(staticFile)
			return
		}
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "404: Not found",
		})
	})

	// Homepage
	router.GET("/", func(c *gin.Context) {
		now := time.Now()
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/overview/%d/%d", now.Year(), now.Month()))
	})

	// Overview
	router.GET("/overview/:year/:month", func(c *gin.Context) {
		year, yearErr := strconv.Atoi(c.Param("year"))
		month, monthErr := strconv.Atoi(c.Param("month"))
		now := time.Now()

		// Date validation
		if yearErr != nil || monthErr != nil || month > 12 || month < 1 {
			c.String(http.StatusBadRequest, "Error: Invalid date")
			return
		} else if year < 2021 {
			c.String(http.StatusBadRequest, "Error: Date must be 2021 onwards")
			return
		} else if year > now.Year() || month > int(now.Month()) {
			c.String(http.StatusBadRequest, "Error: Date must not be in the future")
			return
		}

		// Retrieve summary data from db
		var summary []struct {
			ID         int
			Name       string
			Projected  int
			Actual     int
			Difference int
		}
		db.Raw(read("sql/monthly_summary.sql"), year, month).Scan(&summary)

		// Calculate totals
		var sums [3]int
		for _, v := range summary {
			sums[0] += v.Projected
			sums[1] += v.Actual
			sums[2] += v.Difference
		}

		// Retrieve payments data from db
		var payments []Payment
		db.Joins("Category").Where("year = ? AND month = ?", year, month).Find(&payments)

		// tmp
		income := 2000

		c.HTML(http.StatusOK, "overview.html", gin.H{
			"title":         fmt.Sprintf("Monthly Overview for %s %d", months[month-1], year),
			"payments":      payments,
			"summary":       summary,
			"sums":          sums,
			"projectedNet":  income - sums[0],
			"actualNet":     income - sums[1],
			"netDifference": (income - sums[1]) - (income - sums[0]),
		})
	})

	// Make a payment
	router.GET("/payment", func(c *gin.Context) {
		var categories []Category
		db.Find(&categories)

		c.HTML(http.StatusOK, "payment.html", gin.H{
			"title":      "Submit a Payment",
			"categories": categories,
		})
	})

	// Post point to make a payment
	router.POST("/payment", func(c *gin.Context) {
		c.Request.ParseForm()
		form := c.Request.PostForm

		if !c.Request.PostForm.Has("submit") {
			c.String(http.StatusBadRequest, "Error: Must be submitted via form")
			return
		}

		now := time.Now()

		var payment Payment
		payment.Year = now.Year()
		payment.Month = int(now.Month())
		payment.Day = now.Day()
		payment.Hour = now.Hour()
		payment.Minute = now.Minute()
		payment.CategoryID, _ = strconv.Atoi(form["category"][0])
		payment.Value, _ = strconv.Atoi(form["value"][0])

		db.Create(&payment)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// Categories
	router.GET("/categories", func(c *gin.Context) {
		var categories []Category
		db.Find(&categories)

		c.HTML(http.StatusOK, "categories.html", gin.H{
			"title":      "Categories",
			"categories": categories,
		})
	})

	// Post point to add new category
	router.POST("/category", func(c *gin.Context) {
		c.Request.ParseForm()
		form := c.Request.PostForm

		if !form.Has("submit") {
			c.String(http.StatusBadRequest, "Error: Must be submitted via form")
			return
		}

		var category Category
		category.Name = form["name"][0]
		category.Projected, _ = strconv.Atoi(form["projected_value"][0])

		db.Create(&category)
		c.Redirect(http.StatusMovedPermanently, "/categories")
	})

	// Automatic payments
	router.GET("/automatic", func(c *gin.Context) {
		c.HTML(http.StatusOK, "automatic.html", gin.H{
			"title": "Automatic Payments",
		})
	})

	router.Run("localhost:8080")
}

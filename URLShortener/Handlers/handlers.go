package Handlers

import (
	"math/rand"
	"net/http"
	"time"
	"urlshortener/Database"
	"urlshortener/Models"

	"github.com/labstack/echo/v4"
)


func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}


func CreateShortURL(c echo.Context) error {
	var url Models.URL
	if err := c.Bind(&url); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	url.ShortCode = generateShortCode()

	if err := Database.DB.Create(&url).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert URL into database"})
	}

	return c.JSON(http.StatusOK, url)
}


func RedirectToOriginalURL(c echo.Context) error {
	shortCode := c.Param("shortCode")

	var url Models.URL
	if err := Database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
	}

	return c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
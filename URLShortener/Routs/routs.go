package Routes

import (
	"urlshortener/Handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/shorten", Handlers.CreateShortURL)
	e.GET("/:shortCode", Handlers.RedirectToOriginalURL)
}
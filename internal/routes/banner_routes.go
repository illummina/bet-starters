package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"testBetStarters/internal/handlers"
)

func InitRoutes(e *echo.Echo, db *sqlx.DB) {
	e.GET("/cms/banners", handlers.GetZoneBanner(db))
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"testBetStarters/internal/models"
)

func GetZoneBanner(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		pageID, _ := strconv.Atoi(c.QueryParam("pageId"))
		pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
		sortBy := c.QueryParam("sortBy")
		deviceTypeId := c.QueryParam("deviceTypeId")
		languageCode := c.QueryParam("languageCode")

		var errorMessage string

		rows, err := db.Queryx(
			"EXEC [cms].spGetBanners @DeviceTypeId, @LanguageCode, @PageId, @PageSize, @SortBy, @ErrorMessage OUTPUT",
			sql.Named("DeviceTypeId", deviceTypeId),
			sql.Named("LanguageCode", languageCode),
			sql.Named("PageId", pageID),
			sql.Named("PageSize", pageSize),
			sql.Named("SortBy", sortBy),
			sql.Named("ErrorMessage", sql.Out{Dest: &errorMessage}),
		)
		if err != nil {
			fmt.Println("Error fetching zone_banner:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch zone_banner")
		}
		defer rows.Close()

		if errorMessage != "" {
			return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
		}

		var banners []byte
		for rows.Next() {
			if err = rows.Scan(&banners); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to scan banners: "+err.Error())
			}
		}
		if err = rows.Err(); err != nil {
			return err
		}

		var zoneBanners []models.ZoneBanner
		if err = json.Unmarshal(banners, &zoneBanners); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to unmarshal banners: "+err.Error())
		}

		return c.JSON(http.StatusOK, zoneBanners)
	}
}

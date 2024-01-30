package items

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"openapi/internal/application/stockitem"
	"openapi/internal/domain/stock/item"
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stock"
)

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func Delete(c echo.Context) error {
	// Pre Process
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()
	repository := &item.Repository{Db: db}

	// Validation
	stockitemId := uuid.MustParse(c.Param("stockitemId"))
	if stockitemId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid stock item id")
	}

	found, err := repository.Find(item.Id(stockitemId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock item not found")
	}

	// Main Process
	reqDto := &stockitem.DeleteRequestDto{
		Id: stockitemId,
	}
	_, err = stockitem.Delete(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Post Process
	res := &oapicodegen.OK{}

	return c.JSON(http.StatusOK, res)
}

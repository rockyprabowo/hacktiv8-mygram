package responses

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func WithDeleteSuccess(ctx echo.Context, modelType string) error {
	return ctx.JSON(
		http.StatusOK,
		InfoResult{
			Message: fmt.Sprintf("Your %s has been successfully deleted.", modelType),
		},
	)
}

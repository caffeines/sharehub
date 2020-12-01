package api

import (
	"github.com/caffeines/sharehub/validators"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(endpoint *echo.Group) {
	endpoint.POST("/", register)
}
func register(ctx echo.Context) error {
	_, err := validators.ValidateRegister(ctx)
	if err != nil {
		//
	}
	return nil
}

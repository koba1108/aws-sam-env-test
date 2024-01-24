package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type HealthCheckResponse struct {
	TestEnv1 string `json:"testEnv1"`
	TestEnv2 string `json:"testEnv2"`
	TestEnv3 string `json:"testEnv3"`
}

func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &HealthCheckResponse{
		TestEnv1: os.Getenv("TEST_ENV_1"),
		TestEnv2: os.Getenv("TEST_ENV_2"),
		TestEnv3: os.Getenv("TEST_ENV_3"),
	})
}

type ValidateRequest struct {
	ID   string `param:"id"`   // path parameter
	Name string `query:"name"` // query parameter
	Page int    `query:"page"` // query parameter
}

type ValidateResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Page int    `json:"page,omitempty"`
}

func ValidateHandler(c echo.Context) error {
	var req ValidateRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.ID == "" {
		req.ID = "defaultId"
	}
	if req.Name == "" {
		req.Name = "defaultName"
	}
	if req.Page == 0 {
		req.Page = 1
	}
	return c.JSON(http.StatusOK, &ValidateResponse{
		ID:   req.ID,
		Name: req.Name,
		Page: req.Page,
	})
}

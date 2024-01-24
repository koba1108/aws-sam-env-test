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

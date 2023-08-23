package handlers

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/phramos07/gowebapp/tests"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	successCase := &tests.TestCase{
		TestName: "health test",
		Request: tests.TestRequest{
			Method: http.MethodGet,
			Url:    "/live",
		},
		Response: tests.TestResponse{
			StatusCode: http.StatusOK,
			Body: `{
				"message": "app is running"
			}`,
		},
		HandlerFunc: func(c echo.Context) error {
			return NewHealthHandler().IsAlive(c)
		},
	}

	server := echo.New()
	t.Run(successCase.TestName, func(t *testing.T) {
		mockReq, mockRes := tests.PrepareRequestsForGivenContext(successCase)

		context := server.NewContext(mockReq, mockRes)

		err := successCase.HandlerFunc(context)
		if assert.NoError(t, err) {
			assert.Equal(t, successCase.Response.StatusCode, mockRes.Code)
			assert.JSONEq(t, successCase.Response.Body, mockRes.Body.String())
		}
	})
}

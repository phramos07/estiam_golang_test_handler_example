package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

// this test case might serve for all handlers
type TestCase struct {
	TestName    string
	Request     TestRequest
	Response    TestResponse
	HandlerFunc echo.HandlerFunc
}

type TestRequest struct {
	Method string // GET, PUT, POST, DELETE, etc...
	Url    string // route
	Body   string
}

type TestResponse struct {
	StatusCode int
	Body       string
}

func PrepareRequestsForGivenContext(tcase *TestCase) (*http.Request, *httptest.ResponseRecorder) {
	mockedRequest := httptest.NewRequest(
		tcase.Request.Method,
		tcase.Request.Url,
		strings.NewReader(tcase.Request.Body),
	)
	mockedRequest.Header.Set(
		echo.HeaderContentType,
		echo.MIMEApplicationJSON,
	)

	mockedResponse := httptest.NewRecorder()

	return mockedRequest, mockedResponse
}

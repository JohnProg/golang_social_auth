package integration_tests

import (
	"github.com/golang_social_auth/settings"
	"github.com/golang_social_auth/database"
	"testing"
	"log"
	"os"
	"github.com/golang_social_auth/server"
	"net/http/httptest"
	"fmt"
	"bytes"
	"net/http"
	"github.com/stretchr/testify/assert"
)

const (
	configPath = "../config.json"
)

var (
	testApp   server.Server
)

func TestMain(m *testing.M) {
	testApp = server.Server{}
	if err := settings.Read(configPath); err != nil {
		log.Fatal("Could not read config file at " + configPath + " " + err.Error())
	}

	testApp.Initialize()

	database.DB.Exec("TRUNCATE TABLE users")

	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	testApp.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected response code %d. Got %d\n", expected, actual))
}

func newRequest(t *testing.T, httpStatus int, method, urlStr string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	resp := executeRequest(req)

	checkResponseCode(t, httpStatus, resp.Code)
	return resp
}

func newRequestAuth(t *testing.T, token string, httpStatus int, method, urlStr string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	resp := executeRequest(req)

	checkResponseCode(t, httpStatus, resp.Code)
	return resp
}

package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/lib/middleware"
	"github.com/tegarap/altastore-api/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type adminResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func initEchoTestAPI() *echo.Echo {
	config.InitDbTest("admin")
	e := echo.New()
	return e
}

func InsertData() error {
	admin := models.Admins{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	var err error
	if err = config.Db.Save(&admin).Error; err != nil {
		return err
	}
	return nil
}
func InsertData2() error {
	admin := models.Admins{
		Name:     "Tegar",
		Email:    "tegar@gmail.com",
		Password: "12345678",
	}

	var err error
	if err = config.Db.Save(&admin).Error; err != nil {
		return err
	}
	return nil
}

func TestLoginAdminsController(t *testing.T) {
	var testCases = []struct {
		reqBody      map[string]string
		expectCode   int
		responStatus string
	}{
		{
			reqBody:      map[string]string{"email": "alta@gmail.com", "password": "123"},
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			reqBody:      map[string]string{"email": "kliru@gmail.com", "password": "123123123"},
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertData()
	for _, testCase := range testCases {
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/admin/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, LoginAdminController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response adminResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestRegisterAdminController(t *testing.T) {
	var testCases = []struct {
		reqBody      map[string]string
		expectCode   int
		responStatus string
	}{
		{
			reqBody:      map[string]string{"name": ""},
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			reqBody:      map[string]string{"name": "Test Nama", "email": "iniemail@test.com", "password": "luwakwhitecoffe"},
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			reqBody:      map[string]string{"name": "Test Nama", "email": "iniemail@test.com", "password": "luwakwhitecoffe"},
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	for i, testCase := range testCases {
		if i == 2 {
			config.Db.Migrator().DropTable(&models.Admins{})
		}
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/admin/register", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, RegisterAdminController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response adminResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestGetAdminProfileController(t *testing.T) {
	var testCases = []struct {
		expectCode   int
		responStatus string
	}{
		{
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
		{
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertData()
	InsertData2()
	adminToken, _ := middleware.CreateTokenTest(1, true)
	customerToken, _ := middleware.CreateTokenTest(2, false)
	for i, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/admin/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if i == 0 {
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", customerToken))
		} else {
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", adminToken))
		}
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if i == 2 {
			//drop table
			config.Db.Migrator().DropTable(&models.Admins{})
		}
		err := m.JWT([]byte(os.Getenv("SECRET_JWT_TEST")))(GetAdminProfileController)(context)
		if err != nil {
			return
		}
		assert.Equal(t, testCase.expectCode, res.Code)
		resBody := res.Body.String()

		var response adminResponse
		json.Unmarshal([]byte(resBody), &response)
		assert.Equal(t, testCase.responStatus, response.Status)
	}
}
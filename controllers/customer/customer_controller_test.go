package customer

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

type customerResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func initEchoTestAPI() *echo.Echo {
	config.InitDbTest("customer")
	e := echo.New()
	return e
}

func InsertData() error {
	customer := models.Customers{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	var err error
	if err = config.Db.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}
func InsertData2() error {
	customer := models.Customers{
		Name:     "Tegar",
		Email:    "tegar@gmail.com",
		Password: "12345678",
	}

	var err error
	if err = config.Db.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}

func TestLoginCustomersController(t *testing.T) {
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
		req := httptest.NewRequest(http.MethodPost, "/customers/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, LoginCustomerController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response customerResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestRegisterCustomerController(t *testing.T) {
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
			config.Db.Migrator().DropTable(&models.Customers{})
		}
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/customers/register", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, RegisterCustomerController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response customerResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}

func TestGetCustomerProfileController(t *testing.T) {
	var testCasesCustomer = []struct {
		expectCode   int
		responStatus string
	}{
		{
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	var testCasesAdmin = []struct {
		queryValue   string
		expectCode   int
		responStatus string
	}{
		{
			queryValue:   "2",
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			queryValue:   "99",
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
		},
	}

	e := initEchoTestAPI()
	InsertData()
	InsertData2()
	adminToken, _ := middleware.CreateTokenTest(1, true)
	customerToken, _ := middleware.CreateTokenTest(2, false)
	for _, testCase := range testCasesAdmin {
		req := httptest.NewRequest(http.MethodGet, "/customers/profile", nil)
		q := req.URL.Query()
		q.Add("id", testCase.queryValue)
		req.URL.RawQuery = q.Encode()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", adminToken))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		err := m.JWT([]byte(os.Getenv("SECRET_JWT_TEST")))(GetCustomerProfileController)(context)
		if err != nil {
			return
		}
		assert.Equal(t, testCase.expectCode, res.Code)
		resBody := res.Body.String()

		var response customerResponse
		json.Unmarshal([]byte(resBody), &response)
		assert.Equal(t, testCase.responStatus, response.Status)
	}
	for i, testCase := range testCasesCustomer {
		req := httptest.NewRequest(http.MethodGet, "/customers/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", customerToken))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if i == 1 {
			//drop table
			config.Db.Migrator().DropTable(&models.Customers{})
		}
		err := m.JWT([]byte(os.Getenv("SECRET_JWT_TEST")))(GetCustomerProfileController)(context)
		if err != nil {
			return
		}
		assert.Equal(t, testCase.expectCode, res.Code)
		resBody := res.Body.String()

		var response customerResponse
		json.Unmarshal([]byte(resBody), &response)
		assert.Equal(t, testCase.responStatus, response.Status)
	}
}

func TestGetAllCustomersController(t *testing.T) {
	var testCases = []struct {
		expectCode   int
		responStatus string
	}{
		{
			expectCode:   http.StatusOK,
			responStatus: "success",
		},
		{
			expectCode:   http.StatusBadRequest,
			responStatus: "fail",
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
		if i == 2 {
			config.Db.Migrator().DropTable(&models.Customers{})
		}
		req := httptest.NewRequest(http.MethodGet, "/customers", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if i == 1 {
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", customerToken))
		} else {
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", adminToken))
		}
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		err := m.JWT([]byte(os.Getenv("SECRET_JWT_TEST")))(GetAllCustomersController)(context)
		if err != nil {
			return
		}
		assert.Equal(t, testCase.expectCode, res.Code)
		resBody := res.Body.String()

		var response customerResponse
		json.Unmarshal([]byte(resBody), &response)
		assert.Equal(t, testCase.responStatus, response.Status)
	}
}
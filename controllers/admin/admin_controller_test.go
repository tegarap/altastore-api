package admin

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tegarap/altastore-api/config"
	"github.com/tegarap/altastore-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func initEchoTestAPI() *echo.Echo {
	config.InitDbTest("admin")
	e := echo.New()
	return e
}

func InsertDataAdminForGetAdmins() error {
	user := models.Admins{
		Name:     "Alta",
		Email:    "alta@gmail.com",
		Password: "123",
	}

	var err error
	if err = config.Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
func InsertDataAdminForGetAdmins2() error {
	user := models.Admins{
		Name:     "Tegar",
		Email:    "tegar@gmail.com",
		Password: "12345678",
	}

	var err error
	if err = config.Db.Save(&user).Error; err != nil {
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
	InsertDataAdminForGetAdmins()
	for _, testCase := range testCases {
		requestBody, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/admin/login", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		if assert.NoError(t, LoginAdminController(context)) {
			resBody := res.Body.String()
			assert.Equal(t, testCase.expectCode, res.Code)

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)
			assert.Equal(t, testCase.responStatus, response.Status)
		}
	}
}
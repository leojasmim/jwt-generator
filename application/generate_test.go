package application

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func Test_Generate(t *testing.T) {
	c, r, w := setupServer()
	c.Request = setupRequest()

	r.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, http.StatusOK)
}

func setupServer() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST("/jwt_test", GenerateJWT)
	return c, r, w
}

func setupRequest() *http.Request {
	formData := getFormData()
	request, _ := http.NewRequest(http.MethodPost, "/jwt_test", strings.NewReader(formData))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(formData)))
	return request

}

func getFormData() string {
	data := url.Values{}
	data.Set("payload", `{ "message" : "teste_jwt_generator" }`)
	data.Set("sk", getRSAKey())
	return data.Encode()
}

func getRSAKey() string {
	b, _ := ioutil.ReadFile("../keys/rsa_private_key")
	return string(b)
}

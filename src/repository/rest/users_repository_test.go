package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi (t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "invalid login credentials", "status":"404","error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message": "invalid login credentials", "status":404,"error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": "1", "first_name":"toan2", "last_name": "phung2", "email":"toan22@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshall user login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "http://localhost:8080/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": 1, "first_name":"toan2", "last_name": "phung2", "email":"toan22@gmail.com", "status": "active"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "toan2", user.FirstName)
	assert.EqualValues(t, "phung2", user.LastName)
	assert.EqualValues(t, "toan22@gmail.com", user.Email)
}

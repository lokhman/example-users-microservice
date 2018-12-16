// +build integration

// Run this integration test with the following command:
// $ docker-compose run -e GIN_MODE=test app go test -tags=integration -failfast -v

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lokhman/example-users-microservice/api"
	"github.com/lokhman/example-users-microservice/model"
	"github.com/stretchr/testify/assert"
)

var API *api.API
var Router *gin.Engine

var (
	MockUserInput = api.UserInput{
		Email:     fmt.Sprintf("alex.lokhman.%d@gmail.com", rand.Uint32()),
		Password:  fmt.Sprintf("MyPassword%d", rand.Uint32()),
		FirstName: "Alex",
		LastName:  "Lokhman",
		Nickname:  "VisioN",
		Country:   "RU",
	}
	MockUser = model.User{
		Email:     MockUserInput.Email,
		FirstName: MockUserInput.FirstName,
		LastName:  MockUserInput.LastName,
		Nickname:  MockUserInput.Nickname,
		Country:   MockUserInput.Country,
	}
)

func startup() {
	db := connectDatabase("postgres", os.Getenv("DATABASE_URL"))
	p := connectNSQ(os.Getenv("NSQ_ADDR"))

	API = &api.API{DB: db, NSQ: p}
	Router = createRouter(API)
}

func cleanup() {
	_ = API.DB.Close()
	API.NSQ.Stop()
}

func TestHealthCheck(t *testing.T) {
	startup()
	defer cleanup()

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var out struct {
		Now time.Time `json:"now"`
	}
	dec := json.NewDecoder(w.Body)
	err = dec.Decode(&out)
	assert.Nil(t, err)
	assert.NotZero(t, out.Now)
}

func TestUserCreate(t *testing.T) {
	startup()
	defer cleanup()

	var in = MockUserInput
	data, err := json.Marshal(in)
	assert.Nil(t, err)

	// test for success
	req, err := http.NewRequest("POST", "/users", bytes.NewReader(data))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var out model.User
	dec := json.NewDecoder(w.Body)
	err = dec.Decode(&out)
	assert.Nil(t, err)
	assert.NotZero(t, out.ID)
	assert.Equal(t, MockUser.Email, out.Email)
	assert.Empty(t, out.Password)
	assert.Equal(t, MockUser.FirstName, out.FirstName)
	assert.Equal(t, MockUser.LastName, out.LastName)
	assert.Equal(t, MockUser.Nickname, out.Nickname)
	assert.Equal(t, MockUser.Country, out.Country)
	MockUser.ID = out.ID

	// test for failure
	req, err = http.NewRequest("POST", "/users", bytes.NewReader(data))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	// here we can write more test cases for various scenarios + NSQ publish...
}

func TestUserView(t *testing.T) {
	startup()
	defer cleanup()

	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%d", MockUser.ID), nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var out model.User
	dec := json.NewDecoder(w.Body)
	err = dec.Decode(&out)
	assert.Nil(t, err)
	assert.Equal(t, MockUser, out)
}

func TestUserIndex(t *testing.T) {
	startup()
	defer cleanup()

	req, err := http.NewRequest("GET", "/users?country=RU", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var out []model.User
	dec := json.NewDecoder(w.Body)
	err = dec.Decode(&out)
	assert.Nil(t, err)

	var userFound model.User
	for _, user := range out {
		assert.Equal(t, "RU", user.Country)

		if user.ID == MockUser.ID {
			userFound = user
		}
	}
	assert.Equal(t, MockUser, userFound)
}

func TestUserUpdate(t *testing.T) {
	startup()
	defer cleanup()

	// test if updates
	var in = MockUserInput
	in.Country = "UK"
	data, err := json.Marshal(in)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("/users/%d", MockUser.ID), bytes.NewReader(data))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	// test if is updated
	req, err = http.NewRequest("GET", fmt.Sprintf("/users/%d", MockUser.ID), nil)
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var out model.User
	dec := json.NewDecoder(w.Body)
	err = dec.Decode(&out)
	assert.Nil(t, err)
	assert.Equal(t, in.Country, out.Country)

	// here we can write more test cases for various scenarios + NSQ publish...
}

func TestUserDelete(t *testing.T) {
	startup()
	defer cleanup()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/users/%d", MockUser.ID), nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	// here we can write more test cases for NSQ publish...
}

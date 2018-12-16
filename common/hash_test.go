// +build !integration

// Run this simple unit test with the following command:
// $ docker-compose run app go test -v ./...

package common

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var MockPassword = fmt.Sprintf("MyPassword%d", rand.Uint32())

func TestMustHashPassword(t *testing.T) {
	hash := MustHashPassword(MockPassword)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(MockPassword))
	assert.Nil(t, err)
}

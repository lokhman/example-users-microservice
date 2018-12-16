package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lokhman/example-users-microservice/common"
	"github.com/lokhman/example-users-microservice/model"
	"gopkg.in/go-playground/validator.v8"
)

// @Summary Update user by ID
// @Accept  json
// @Produce json
// @Param   id path int true "User ID" mininum(1)
// @Param   user body api.UserInput true "New user details"
// @Success 204 ""
// @Failure 400 {object} common.HTTPError
// @Failure 404 {object} common.HTTPError
// @Failure 422 {object} common.HTTPError
// @Router  /users/{id} [put]
func (api *API) UserUpdateHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, invalidUserIDError)
		return
	}

	var user model.User
	if err = api.DB.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, userNotFoundError)
			return
		}
		panic(err)
	}

	var in UserInput

	// see `api.UserCreateHandler` for more details
	if err := c.ShouldBindJSON(&in); err != nil {
		code := http.StatusBadRequest
		if _, ok := err.(validator.ValidationErrors); ok {
			code = http.StatusUnprocessableEntity
		}
		c.JSON(code, common.HTTPError{Err: err.Error()})
		return
	}

	user.Email = in.Email
	user.FirstName = in.FirstName
	user.LastName = in.LastName
	user.Nickname = in.Nickname
	user.Country = in.Country

	// for password change I'd rather introduce a separate endpoint or transform request method to PATCH but with care,
	// as later may face problems with nullable fields (pointer type), since identifying if request body contains
	// a specific field or the field is empty will be tricky
	//
	// can also allow empty password in request body (as an exception) and update it only if provided
	// the implementation entirely depends on the project design and some limitations of Go and framework
	user.Password = common.MustHashPassword(in.Password)

	// try to save user entity to the database
	if err = api.DB.Save(&user).Error; err != nil {
		if common.IsUniqueConstraintError(err, model.UserEmailUniqueConstraintName) {
			c.JSON(http.StatusUnprocessableEntity, common.HTTPError{
				Err: fmt.Sprintf(`User with email "%s" exists`, in.Email),
			})
			return
		}
		panic(err)
	}

	// try to publish message to the queue under "user.update" topic
	if err := common.NSQPublish(api.NSQ, "user.update", user); err != nil {
		// correct behaviour should be defined by requirements and may include:
		// - rollback transaction above and panic
		// - log error with data and continue
		// - log error and continue
		// - panic! (yes, I'm cutting corners...)
		panic(err)
	}
	// some meaningful logs to default logger
	log.Printf("[users] user with ID %d was updated", user.ID)

	c.JSON(http.StatusNoContent, nil)
}

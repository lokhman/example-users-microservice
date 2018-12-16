package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lokhman/example-users-microservice/common"
	"github.com/lokhman/example-users-microservice/model"
	"gopkg.in/go-playground/validator.v8"
)

// User input structure.
// We may reuse `model.User` but may lead to various complications,
// e.g. `ID` field needs to be emptied before INSERT/UPDATE, etc.
type UserInput struct {
	Email     string `json:"email" binding:"required,email" example:"alex.lokhman@gmail.com"`
	Password  string `json:"password" binding:"required,min=3,max=72" example:"MyPassword"`
	FirstName string `json:"first_name" binding:"required,max=72" example:"Alex"`
	LastName  string `json:"last_name" binding:"required,max=72" example:"Lokhman"`
	Nickname  string `json:"nickname" binding:"required,max=32" example:"VisioN"`
	Country   string `json:"country" binding:"required,len=2,alpha" example:"RU"`
}

// @Summary Create new user
// @Accept  json
// @Produce json
// @Param   user body api.UserInput true "New user details"
// @Success 200 {object} model.User
// @Failure 400 {object} common.HTTPError
// @Failure 422 {object} common.HTTPError
// @Router  /users [post]
func (api *API) UserCreateHandler(c *gin.Context) {
	var in UserInput

	// we support only `application/json` content type but can be extended
	// with `c.ShouldBind()` that takes request content type into account
	// this will require more type specific description in input structures
	if err := c.ShouldBindJSON(&in); err != nil {
		code := http.StatusBadRequest
		if _, ok := err.(validator.ValidationErrors); ok {
			code = http.StatusUnprocessableEntity
		}
		c.JSON(code, common.HTTPError{Err: err.Error()})
		return
	}

	// new entity
	user := model.User{
		Email:     in.Email,
		Password:  common.MustHashPassword(in.Password),
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Nickname:  in.Nickname,
		Country:   in.Country,
	}

	// try to save user entity to the database
	if err := api.DB.Create(&user).Error; err != nil {
		if common.IsUniqueConstraintError(err, model.UserEmailUniqueConstraintName) {
			c.JSON(http.StatusUnprocessableEntity, common.HTTPError{
				Err: fmt.Sprintf(`User with email "%s" exists`, in.Email),
			})
			return
		}
		panic(err)
	}

	// try to publish message to the queue under "user.create" topic
	if err := common.NSQPublish(api.NSQ, "user.create", user); err != nil {
		// correct behaviour should be defined by requirements and may include:
		// - rollback transaction above and panic
		// - log error with data and continue
		// - log error and continue
		// - panic! (yes, I'm cutting corners...)
		panic(err)
	}

	// some meaningful logs to default logger
	log.Printf("[users] user with ID %d was created", user.ID)

	c.JSON(http.StatusOK, user)
}

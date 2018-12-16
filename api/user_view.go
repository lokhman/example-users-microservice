package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lokhman/example-users-microservice/common"
	"github.com/lokhman/example-users-microservice/model"
)

var (
	invalidUserIDError = common.HTTPError{Err: "Invalid user ID"}
	userNotFoundError  = common.HTTPError{Err: "User cannot be found"}
)

// @Summary View user details
// @Accept  json
// @Produce json
// @Param   id path int true "User ID" mininum(1)
// @Success 200 {object} model.User
// @Failure 404 {object} common.HTTPError
// @Router  /users/{id} [get]
func (api *API) UserViewHandler(c *gin.Context) {
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

	c.JSON(http.StatusOK, user)
}

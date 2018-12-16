package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lokhman/example-users-microservice/common"
	"github.com/lokhman/example-users-microservice/model"
)

// @Summary Delete user by ID
// @Accept  json
// @Produce json
// @Param   id path int true "User ID" mininum(1)
// @Success 204 ""
// @Failure 404 {object} common.HTTPError
// @Router  /users/{id} [delete]
func (api *API) UserDeleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, invalidUserIDError)
		return
	}

	var user model.User
	if err = api.DB.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// DELETE request is idempotent, so we may show that request was successful
			c.JSON(http.StatusNoContent, nil)
			return
		}
		panic(err)
	}

	// try to delete user entity from the database
	if err = api.DB.Delete(&user).Error; err != nil {
		panic(err)
	}

	// try to publish message to the queue under "user.delete" topic
	if err := common.NSQPublish(api.NSQ, "user.delete", user); err != nil {
		// correct behaviour should be defined by requirements and may include:
		// - rollback transaction above and panic
		// - log error with data and continue
		// - log error and continue
		// - panic! (yes, I'm cutting corners...)
		panic(err)
	}
	// some meaningful logs to default logger
	log.Printf("[users] user with ID %d was deleted", user.ID)

	c.JSON(http.StatusNoContent, nil)
}

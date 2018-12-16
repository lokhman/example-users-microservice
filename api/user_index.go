package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lokhman/example-users-microservice/model"
)

// @Summary List users
// @Accept  json
// @Produce json
// @Param   country query string false "User country" minlength(2) maxlength(2)
// @Success 200 {array} model.User
// @Router  /users [get]
func (api *API) UserIndexHandler(c *gin.Context) {
	var users []model.User

	db := api.DB
	if country, ok := c.GetQuery("country"); ok {
		db = db.Where(&model.User{Country: country})
	}

	// PostgreSQL doesn't have default order by primary key
	// (entities in the list do not "shuffle" when we update one)
	if err := db.Order("id").Find(&users).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, users)
}

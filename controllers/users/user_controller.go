package users

import (
	"database/sql"
	"jwt_auth/domain/users"
	"jwt_auth/services"
	"jwt_auth/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var rawUser users.User
		if err := c.ShouldBindJSON(&rawUser); err != nil {
			errMsg := errors.NewBadRequestError("json struct in post body is not valid")
			c.AbortWithStatusJSON(errMsg.Status, errMsg)
			return
		}
		user, err := services.CreateUser(rawUser, db)
		if err != nil {
			c.AbortWithStatusJSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

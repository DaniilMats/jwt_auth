package cmd

import (
	"database/sql"
	"jwt_auth/controllers/users"

	"github.com/gin-gonic/gin"
)

func mapURLs(router *gin.Engine, db *sql.DB) {
	v1 := router.Group("api/v1")
	{
		v1.POST("/register", users.Register(db))
	}
}

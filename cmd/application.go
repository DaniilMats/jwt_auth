package cmd

import (
	"jwt_auth/config"
	"jwt_auth/datasource/mysql/user_db"

	"github.com/gin-gonic/gin"
)

func RunApplication(c *config.Config) {
	db, err := user_db.InitConnectionToDB(c)
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	mapURLs(router, db)
	err = router.Run(":8081")
	if err != nil {
		panic(err)
	}
}

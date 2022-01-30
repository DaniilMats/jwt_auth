package user_db

import (
	"database/sql"
	"fmt"
	"jwt_auth/config"

	_ "github.com/go-sql-driver/mysql"
)

func InitConnectionToDB(c *config.Config) (*sql.DB, error) {
	dataSourceHost := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DbUserName, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
	db, err := sql.Open("mysql", dataSourceHost)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

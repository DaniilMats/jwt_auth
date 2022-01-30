package users

import (
	"database/sql"
	"jwt_auth/utils/errors"
)

const insertQuery = `INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);`

func Save(user *User, db *sql.DB) *errors.RestErr {
	statement, err := db.Prepare(insertQuery)
	defer statement.Close()
	if err != nil {
		return errors.NewInternalServiceError("was not able to prepare sql statement")
	}
	result, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return errors.NewInternalServiceError("was not able to save user to DB")
	}
	userID, err := result.LastInsertId()
	user.ID = userID
	return nil
}

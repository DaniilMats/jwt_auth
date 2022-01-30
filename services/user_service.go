package services

import (
	"database/sql"
	"jwt_auth/domain/users"
	"jwt_auth/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User, db *sql.DB) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	encrypred, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewInternalServiceError(err.Error())
	}
	user.Password = string(encrypred)
	if saveError := users.Save(&user, db); err != nil {
		return nil, saveError
	}

	return &user, nil
}

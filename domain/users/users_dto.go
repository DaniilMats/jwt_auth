package users

import (
	"jwt_auth/utils/errors"
	"strings"
)

type User struct {
	ID        int64  `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}

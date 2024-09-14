package user

import "go_starter/pkg/entity"

type User struct {
	entity.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}

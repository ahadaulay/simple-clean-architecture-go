package userrepository

import "belajar-clean/models/dto"

type UserRepository interface {
	GetAllUser() ([]dto.User , error)
	CreateUser(input dto.UserCreate) (dto.User, error)
}

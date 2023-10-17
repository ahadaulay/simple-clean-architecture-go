package userservice

import (
	"belajar-clean/models/dto"
	userrepository "belajar-clean/repositories/userRepository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUser() ([]dto.User,error) 
	CreateUser(input dto.UserCreate) (dto.User,error)
}

type userImplementation struct{
	repository userrepository.UserRepository
}

func (ui *userImplementation) GetAllUser() ([]dto.User, error)   {

	data, err := ui.repository.GetAllUser()

	if err != nil {
		return []dto.User{} , err
	}

	return data , nil
}

func (ui *userImplementation) CreateUser(input dto.UserCreate) (dto.User , error)  {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	
	if err != nil{
		return dto.User{}, err
	}
	input.Password = string(hashPassword)
	data, err := ui.repository.CreateUser(input)

	if err!= nil {
		return dto.User{} , err
	}

	return data, nil
}

func NewUserService(repository userrepository.UserRepository) UserService  {
	return &userImplementation{
		repository: repository,
	}
}
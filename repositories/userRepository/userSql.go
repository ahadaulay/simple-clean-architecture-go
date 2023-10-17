package userrepository

import (
	"belajar-clean/models/dto"
	"belajar-clean/models/model"

	"gorm.io/gorm"
)

type userImplementation struct {
	db *gorm.DB
}

func (ui *userImplementation) GetAllUser() ([]dto.User,error)  {
	var userModel []model.User
	var UserDto []dto.User

	err := ui.db.Find(&userModel).Error
	
	if err != nil {
		return []dto.User{} , err
	}

	for _, data := range userModel {
		UserDto = append(UserDto, dto.User(data))
	}

	return UserDto , nil
}

func (ui *userImplementation) CreateUser(input dto.UserCreate) (dto.User , error)  {
	var userCreate model.User = model.User{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}

	var userDto model.User = model.User{}

	var result dto.User

	err := ui.db.Create(&userCreate).Error
	

	ui.db.Last(&userDto)

	result = dto.User(userDto)

	if err != nil {
		return dto.User{} , err
	}

	return result , nil

}

func NewUserRepository(db *gorm.DB) UserRepository  {
	return &userImplementation{
		db : db,
	}
}
package userService

import (
	"errors"

	"github.com/google/uuid"
	"github.com/waxer59/basic-go-fiber-api/database"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"
)

func GetUserById(id string) (*userModels.User, error) {
	db := database.DB

	var user userModels.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return nil, errors.New("No user found")
	}

	return &user, nil
}

func GetUserByEmail(email string) (*userModels.User, error) {
	db := database.DB

	var user userModels.User

	db.Find(&user, "email = ?", email)

	if user.ID == uuid.Nil {
		return nil, errors.New("No user found")
	}

	return &user, nil
}

func CreateUser(user userModels.User) (*userModels.User, error) {
	db := database.DB

	err := db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id string, updateUser userModels.UpdateUser) (*userModels.User, error) {
	db := database.DB

	user, err := GetUserById(id)

	if err != nil {
		return nil, err
	}

	if err := updateUser.ValidateFields(); err != nil {
		return nil, err
	}

	user.Name = updateUser.Name
	user.Email = updateUser.Email

	if err := updateUser.HashPassword(); err != nil {
		return nil, err
	}

	db.Model(user).Updates(&updateUser)

	return user, nil
}

func DeleteUser(id string) (*userModels.User, error) {
	db := database.DB

	user, err := GetUserById(id)

	if err != nil {
		return nil, err
	}

	err = db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsers() []userModels.User {
	db := database.DB

	var users []userModels.User

	db.Find(&users)

	return users
}

package authService

import (
	"github.com/waxer59/basic-go-fiber-api/internal/user/userService"
	"github.com/waxer59/basic-go-fiber-api/internal/utils/bcryptUtils"
	"github.com/waxer59/basic-go-fiber-api/internal/utils/jwtUtils"
)

func UserLogin(email string, password string) (string, error) {
	user, err := userService.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcryptUtils.Compare(password, user.Password)

	if err != nil {
		return "", err
	}

	token, err := jwtUtils.NewJwt(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

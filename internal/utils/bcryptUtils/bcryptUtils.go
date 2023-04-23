package bcryptUtils

import "golang.org/x/crypto/bcrypt"

func Encript(password string) (string, error) {
	encriptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(encriptedPassword), nil
}

func Compare(password string, encriptedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encriptedPassword), []byte(password))
}

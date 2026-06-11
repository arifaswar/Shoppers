package service

import (
	"golang.org/x/crypto/bcrypt"
	"shoppers/internal/domain"
	"shoppers/internal/repository"
)

func Register(
	name string,
	email string,
	password string,
	role string,
) error {
	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := domain.User{
		Name: name,
		Email: email,
		Password: string(hashPassword),
		Role: role,
	}
	return repository.CreateUser(&user)
}

func Login(
	email string,
	password string,
) (string, error) {

	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", err
	}

	token, err := GenerateToken(
		user.ID.String(),
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
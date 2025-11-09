package service

import (
	"backend-app/internal/domain"
	"backend-app/internal/dto"
)

type UserService struct {
}

func (userService UserService) FindUserByEmail(email string) (*domain.User, error) {

	return nil, nil
}

func (userService UserService) SignupUser(userPayload dto.UserSignupRequestDto) (string, error) {

	return "find some dummy token", nil
}

func (userService UserService) Login(userPayload any) (string, error) {

	return "user logged in", nil
}

func (userService UserService) GetVerificationCode(e *domain.User) (int, error) {

	return 0, nil
}

func (userService UserService) VerifyCode(email string, code int) (bool, error) {

	return true, nil
}

func (userService UserService) GetUserProfile(userID uint) (*domain.User, error) {

	return nil, nil
}

func (userService UserService) AddItemToCart(userID uint, itemID uint) error {

	return nil
}

func (userService UserService) GetCartItems(userID uint) ([]any, error) {
	return nil, nil
}

func (userService UserService) ClearCart(userID uint) error {

	return nil
}

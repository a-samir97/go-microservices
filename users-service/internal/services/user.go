package services

import (
	"log"
	"users/api/dto"
	"users/internal/domain"
	"users/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(user domain.User) (*domain.User, error) {

	return us.repo.Save(&user)
}

func (us *UserService) FindByID(id int) (*domain.User, error) {
	return us.repo.FindByID(id)
}

func (us *UserService) UpdateUser(updatedUser domain.User, id int) (*domain.User, error) {

	return us.repo.Update(updatedUser, id)
}

func (us *UserService) DeleteUser(id string) error {

	return us.repo.Delete(id)
}

func (us *UserService) Login(email, password string) (*dto.LoginResponse, error) {
	// check email
	user, err := us.repo.FindByEmail(email)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	loginResponse := &dto.LoginResponse{
		Email: user.Email,
		Token: tokenString,
	}

	return loginResponse, nil
}

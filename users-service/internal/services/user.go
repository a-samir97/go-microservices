package services

import (
	"log"
	"time"
	"users/api/dto"
	"users/internal/domain"
	"users/internal/event"
	"users/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      repository.UserRepository
	publisher event.Publisher
}

func NewUserService(repo repository.UserRepository, publisher event.Publisher) *UserService {
	return &UserService{repo: repo, publisher: publisher}
}

func (us *UserService) CreateUser(user domain.User) (*domain.User, error) {

	u, err := us.repo.Save(&user)
	if err != nil {
		return nil, err
	}
	// publish event
	userCreatedEvent := event.UserCreatedEvent{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: time.Now(),
	}

	err = us.publisher.PublishUserCreatedEvent(userCreatedEvent)
	if err != nil {
		log.Println(err)
	}

	return u, nil
}

func (us *UserService) FindByID(id int) (*domain.User, error) {
	return us.repo.FindByID(id)
}

func (us *UserService) UpdateUser(updatedUser domain.User, id int) (*domain.User, error) {
	user, err := us.repo.Update(updatedUser, id)

	// publish event
	userUpdatedEvent := event.UserUpdatedEvent{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		UpdatedAt: time.Now(),
	}

	err = us.publisher.PublishUserUpdatedEvent(userUpdatedEvent)
	if err != nil {
		log.Println(err)
	}
	
	return user, err
}

func (us *UserService) DeleteUser(id string) error {
	err := us.repo.Delete(id)

	log.Println(err)

	// publish event
	userDeletedEvent := event.UserDeletedEvent{
		ID: id,
	}

	err = us.publisher.PublishUserDeletedEvent(userDeletedEvent)
	if err != nil {
		log.Println(err)
	}

	return err
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

package usecases

import (
	"errors"
	"users-crud/internal/dto"
	"users-crud/internal/models"
	"users-crud/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(input dto.RegisterInput) (*models.User, error)
	Login(input dto.LoginInput) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, input dto.UpdateUserInput) (*models.User, error)
	DeleteUser(id uint) error
	UpdateUserRole(id uint, newRole string) (*models.User, error)
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) RegisterUser(input dto.RegisterInput) (*models.User, error) {
	existing, _ := u.repo.FindByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("email já cadastrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
		Role:     "user",
	}

	err = u.repo.Create(user)
	return user, err
}

func (u *userUsecase) Login(input dto.LoginInput) (*models.User, error) {
	user, err := u.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("email ou senha inválidos")
	}

	return user, nil
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.FindAll()
}

func (u *userUsecase) GetUserByID(id uint) (*models.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) UpdateUser(id uint, input dto.UpdateUserInput) (*models.User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email

	err = u.repo.Update(user)
	return user, err
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}

func (u *userUsecase) UpdateUserRole(id uint, newRole string) (*models.User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Role = newRole
	err = u.repo.Update(user)
	return user, err
}

package services

import (
	"errors"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/dto"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/models"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/repository"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/utils"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func (a *authService) Register(req dto.RegisterRequest) error {
	_, err := a.userRepo.GetByEmail(req.Email)
	if err == nil {
		return errors.New("Email already exists")
	}
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil{
		return err
	}
	user := &models.User{
		Name: req.Name,
		Email: req.Email,
		Password: hashPassword,
	}
	return a.userRepo.Create(user)
}

func (a *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userRepo.GetByEmail(req.Email)
	if err != nil{
		return nil, errors.New("Invalid Credentials")
	}
	if !utils.CheckPassword(req.Password, user.Password){
		return nil, errors.New("Invalid Credentials")
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil{
		return nil, err
	}
	return &dto.LoginResponse{Token: token}, nil
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

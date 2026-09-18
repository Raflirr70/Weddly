package usecase

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Raflirr70/Weddly/internal/user/entity"
	"github.com/Raflirr70/Weddly/internal/user/repository"
	"github.com/Raflirr70/Weddly/pkg/apperror"
)

type UserUsecase struct {
	repo *repository.UserRepository
}

func NewUserUsecase(repo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) List() ([]entity.UserResponse, *apperror.AppError) {
	users, err := u.repo.FindAll()
	if err != nil {
		return nil, apperror.Internal("Failed to list users")
	}

	result := make([]entity.UserResponse, 0, len(users))
	for i := range users {
		result = append(result, entity.ToUserResponse(&users[i]))
	}
	return result, nil
}

func (u *UserUsecase) Create(req entity.CreateUserRequest) (*entity.UserResponse, *apperror.AppError) {
	if req.Role == "" {
		req.Role = "admin"
	}

	if _, err := u.repo.FindByUsername(req.Username); err == nil {
		return nil, apperror.BadRequest("Username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperror.Internal("Failed to hash password")
	}

	user := entity.User{
		Username: req.Username,
		Password: string(hash),
		Role:     req.Role,
		Status:   true,
	}
	if err := u.repo.Create(&user); err != nil {
		return nil, apperror.Internal("Failed to create user")
	}

	result := entity.ToUserResponse(&user)
	return &result, nil
}

func (u *UserUsecase) Update(id uint, req entity.UpdateUserRequest) (*entity.UserResponse, *apperror.AppError) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, apperror.NotFound("User Not Found")
	}

	if req.Username != "" && req.Username != user.Username {
		if existing, err := u.repo.FindByUsername(req.Username); err == nil && existing.ID != id {
			return nil, apperror.BadRequest("Username already exists")
		}
		user.Username = req.Username
	}

	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, apperror.Internal("Failed to hash password")
		}
		user.Password = string(hash)
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	if err := u.repo.Update(user); err != nil {
		return nil, apperror.Internal("Failed to update user")
	}

	result := entity.ToUserResponse(user)
	return &result, nil
}

func (u *UserUsecase) Delete(id uint) *apperror.AppError {
	if _, err := u.repo.FindByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("User Not Found")
		}
		return apperror.Internal("Failed to find user")
	}

	if err := u.repo.Delete(id); err != nil {
		return apperror.Internal("Failed to delete user")
	}
	return nil
}
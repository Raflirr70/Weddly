package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Raflirr70/Weddly/internal/auth/entity"
	"github.com/Raflirr70/Weddly/internal/auth/repository"
	"github.com/Raflirr70/Weddly/pkg/apperror"
)

const tokenTTL = 24 * time.Hour

type AuthUsecase struct {
	repo      *repository.AuthRepository
	jwtSecret []byte
}

func NewAuthUsecase(repo *repository.AuthRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{repo: repo, jwtSecret: []byte(jwtSecret)}
}

func (u *AuthUsecase) Login(username, password string) (*entity.LoginResponse, *apperror.AppError) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return nil, apperror.NotFound("User Not Found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, apperror.Unauthorized("Invalid username or password")
	}

	token, err := u.generateToken(&entity.Claims{
		UserID: user.ID,
		Role:   user.Role,
	})
	if err != nil {
		return nil, apperror.Internal("Failed to generate token")
	}

	if err := u.repo.SaveToken(user.ID, token, tokenTTL); err != nil {
		return nil, apperror.Internal("Failed to save token")
	}

	return &entity.LoginResponse{
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}, nil
}

func (u *AuthUsecase) generateToken(claims *entity.Claims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.jwtSecret)
}
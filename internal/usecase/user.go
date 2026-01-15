package usecase

import (
	"clean-architecture-practice/internal/domain/user"
	"context"
	"errors"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, id, name, email string, age int) error
	GetUser(ctx context.Context, id string) (*user.User, error)
}

type userUsecase struct {
	repo user.Repository
}

func NewUserUsecase(repo user.Repository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, id, name, email string, age int) error {
	// ここでドメインルールのバリデーションなどを行うことができます
	if id == "" || name == "" {
		return errors.New("invalid input")
	}

	newUser := &user.User{
		ID:    id,
		Name:  name,
		Email: email,
		Age: age,
	}

	return u.repo.Save(ctx, newUser)
}

func (u *userUsecase) GetUser(ctx context.Context, id string) (*user.User, error) {
	return u.repo.FindByID(ctx, id)
}

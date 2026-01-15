package repository

import (
	"clean-architecture-practice/internal/domain/user"
	"context"
	"errors"
	"sync"
)

type userMemoryRepository struct {
	mu    sync.RWMutex
	users map[string]*user.User
}

func NewUserMemoryRepository() user.Repository {
	return &userMemoryRepository{
		users: make(map[string]*user.User),
	}
}

func (r *userMemoryRepository) Save(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[u.ID] = u
	return nil
}

func (r *userMemoryRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

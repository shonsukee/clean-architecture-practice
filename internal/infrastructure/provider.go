package infrastructure

import (
    "clean-architecture-practice/internal/infrastructure/repository"
    "clean-architecture-practice/internal/interfaces/handler"
    "clean-architecture-practice/internal/usecase"
    "github.com/google/wire"
)

// RepositoryProviders はリポジトリの依存関係セット
var RepositoryProviders = wire.NewSet(
    repository.NewUserMemoryRepository,
)

// UsecaseProviders はユースケースの依存関係セット
var UsecaseProviders = wire.NewSet(
    usecase.NewUserUsecase,
)

// HandlerProviders はハンドラーの依存関係セット
var HandlerProviders = wire.NewSet(
    handler.NewUserHandler,
)

// AllProviders はすべてのプロバイダーを含む
var AllProviders = wire.NewSet(
    RepositoryProviders,
    UsecaseProviders,
    HandlerProviders,
)
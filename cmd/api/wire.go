//go:build wireinject
// +build wireinject

package main

import (
	"clean-architecture-practice/internal/infrastructure"
	"clean-architecture-practice/internal/interfaces/handler"

	"github.com/google/wire"
)

func Wire() (*handler.UserHandler, error) {
	wire.Build(infrastructure.AllProviders)
	return nil, nil
}

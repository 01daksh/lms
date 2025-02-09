//go:build wireinject
// +build wireinject

package book

import (
	"lms/internal/book/handlers"
	"lms/internal/book/repository"
	"lms/internal/book/service"

	"github.com/google/wire"
)

// Wire Set containing all dependencies
var BookSet = wire.NewSet(
	repository.NewBookIssuanceRepo,
	service.NewBookIssuanceService,
	handlers.NewBookIssuanceHandler,
)

// Generate dependency injection function (only defined here for Wire)
func InitializeBookHandler() handlers.BookIssuanceHandler {
	wire.Build(BookSet)
	return nil // Wire will replace this with actual dependency injection logic
}

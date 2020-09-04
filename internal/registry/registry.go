package registry

import (
	"github.com/nhannvt/resume/configs"
	"github.com/nhannvt/resume/internal/domain/validation"
	"github.com/nhannvt/resume/internal/infrastracture"
	"github.com/nhannvt/resume/internal/interface/handler"
	"github.com/nhannvt/resume/internal/usecase"
	"github.com/nhannvt/resume/pkg/http_client"
	_ "github.com/nhannvt/resume/pkg/log"
	"github.com/nhannvt/resume/pkg/redis_client"
)

// Registry represents service locater to resolve depencency during services.
type Registry interface {
	NewHealthCheckHandler() handler.HealthCheckHandler
	NewNotFoundHandler() handler.NotFoundHandler
	NewHealthCheckUseCase() usecase.HealthCheckUseCase

	NewHTTPClient() *http_client.HTTPClient
	NewRedisClient() *redis_client.RedisClient
	NewValidator() validation.Validator
}

type registry struct {
	config configs.Config
}

func NewRegistry(config configs.Config) Registry {
	return &registry{config}
}

func (repo *registry) NewHealthCheckHandler() handler.HealthCheckHandler {
	return handler.NewHealthCheckHandler(repo.NewHealthCheckUseCase())
}

func (repo *registry) NewNotFoundHandler() handler.NotFoundHandler {
	return handler.NewNotFoundHandler()
}

func (repo *registry) NewHealthCheckUseCase() usecase.HealthCheckUseCase {
	return usecase.NewHealthCheckUseCase()
}

func (repo *registry) NewHTTPClient() *http_client.HTTPClient {
	return http_client.NewHTTPClient("https://4rum.vn", "go-client", "123131")
}

func (repo *registry) NewRedisClient() *redis_client.RedisClient {
	return redis_client.NewRedisClient("redis:6379", 0, "")
}

func (repo *registry) NewValidator() validation.Validator {
	return infrastracture.NewValidator(validation.ValidationStructs)
}

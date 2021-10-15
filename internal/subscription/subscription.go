package subscription

import (
	"Gymondo/internal/logger"
	"Gymondo/internal/metrics"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	//Update(ctx context.Context, queueName string)
}

type service struct {
	validate   *validator.Validate
	mysql      MySQLRepository
	redis      RedisRepository
	logger     *logger.StandardLogger
	prometheus *metrics.Prometheus
	config     *Config
}

func CreateService(
	config *Config,
	logger *logger.StandardLogger,
	mysql MySQLRepository,
	redis RedisRepository,
	prometheus *metrics.Prometheus,
	validator *validator.Validate) Service {
	return &service{
		validate:   validator,
		redis:      redis,
		mysql:      mysql,
		logger:     logger,
		prometheus: prometheus,
		config:     config,
	}
}

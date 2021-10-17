package subscription

import (
	"Gymondo/internal/logger"
	"Gymondo/internal/metrics"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	UserGenerator(count int)
	ProductGenerator(count int)
	PlanGenerator()
	GetProductsList() ([]Product, error)
	GetProductById(id int) (Product, error)
	BuyProduct(bpr BuyRequest) (UserPlan, error)
	FetchPlansByUserId(userId int) ([]UserPlan, error)
}

type service struct {
	validate   *validator.Validate
	psql       PgSQLRepository
	redis      RedisRepository
	logger     *logger.StandardLogger
	prometheus *metrics.Prometheus
	config     *Config
}

func CreateService(
	config *Config,
	logger *logger.StandardLogger,
	psql PgSQLRepository,
	redis RedisRepository,
	prometheus *metrics.Prometheus,
	validator *validator.Validate) Service {
	return &service{
		validate:   validator,
		redis:      redis,
		psql:       psql,
		logger:     logger,
		prometheus: prometheus,
		config:     config,
	}
}

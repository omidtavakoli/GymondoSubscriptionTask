package subscription

type RedisRepository interface {
	IsReady() bool
}

type ExternalDriver interface {
}

type PgSQLRepository interface {
	CreateUser(email, username, fullname string) error
	CreateProduct(name string) error
	GetUserByEmail(email string) (User, error)
	GetProducts() ([]Product, error)
}

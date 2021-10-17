package subscription

type RedisRepository interface {
	IsReady() bool
}

type ExternalDriver interface {
}

type PgSQLRepository interface {
	CreateUser(email, username, fullname string) error
	CreateProduct(name string) error
	CreatePlan(name string, price, discount, durationDays int, product Product) (uint64, error)
	GetUserByEmail(email string) (User, error)
	GetProducts() ([]Product, error)
	GetProduct(id int) (Product, error)
	BuyProduct(bpr BuyRequest) (UserPlan, error)
}

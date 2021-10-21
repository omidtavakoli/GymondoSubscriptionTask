package subscription

type RedisRepository interface {
	IsReady() bool
}

type ExternalDriver interface {
}

type PgSQLRepository interface {
	CreateUser(email, username, fullname string) (User, error)
	CreateProduct(name string) error
	CreateVoucher(cvr CreateVoucherRequest) (Voucher, error)
	CreatePlan(name string, price, discount, durationDays int, product Product) (uint64, error)
	CreateVoucherPlan(plan Plan, voucher Voucher) (uint64, error)
	GetUserByEmail(email string) (User, error)
	GetProducts() ([]Product, error)
	GetPlans() ([]Plan, error)
	GetVouchers() ([]Voucher, error)
	GetProduct(id int) (Product, error)
	BuyProduct(bpr BuyRequest) (UserPlan, error)
	FetchPlansByUserId(userId int) ([]Status, error)
	ChangeUserPlanStatus(status ChangeStatus) error
}

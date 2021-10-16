package subscription

type RedisRepository interface {
	IsReady() bool
}

type ExternalDriver interface {
}

type PgSQLRepository interface {
	GetUser(email string) error
}

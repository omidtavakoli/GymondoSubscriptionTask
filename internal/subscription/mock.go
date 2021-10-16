package subscription

import "github.com/brianvoe/gofakeit"

func (s service) UserGenerator(count int) {
	for i := 0; i < count; i++ {
		email := gofakeit.Email()
		username := gofakeit.BeerName()
		fullName := gofakeit.Name()
		err := s.psql.CreateUser(email, username, fullName)
		if err != nil {
			s.logger.Errorf("err creating user:%s", err)
		}
	}
}

func (s service) ProductGenerator(count int) {
	for i := 0; i < count; i++ {
		name := gofakeit.DomainName()
		err := s.psql.CreateProduct(name)
		if err != nil {
			s.logger.Errorf("err creating product:%s", err)
		}
	}
}

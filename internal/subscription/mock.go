package subscription

import "github.com/brianvoe/gofakeit"

func (s service) UserGenerator(count int) {
	for i := 0; i < count; i++ {
		email := gofakeit.Email()
		username := gofakeit.BeerName()
		fullName := gofakeit.Name()
		cu, err := s.psql.CreateUser(email, username, fullName)
		if err != nil {
			s.logger.Errorf("err creating user:%s", err)
		} else {
			s.logger.Infof("User:%s created", cu.UserName)
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

func (s service) PlanGenerator() {
	products, err := s.GetProductsList()
	if err != nil {
		s.logger.Errorf("err fetching products:%s", err)
	} else {
		for i, product := range products {
			plan, cpErr := s.psql.CreatePlan("LifeTime", (i+1)*1000, 10, -1, product)
			if cpErr != nil {
				s.logger.Errorf("err creating product:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", plan)
			}

			oneMonthPlan, cpErr := s.psql.CreatePlan("One Month", (i+1)*100, 2, 30, product)
			if cpErr != nil {
				s.logger.Errorf("err creating product:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", oneMonthPlan)
			}

			threeMonthPlan, cpErr := s.psql.CreatePlan("Three Months", (i+1)*300, 6, 90, product)
			if cpErr != nil {
				s.logger.Errorf("err creating product:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", threeMonthPlan)
			}

			sixMonthPlan, cpErr := s.psql.CreatePlan("Six Months", (i+1)*300, 8, 180, product)
			if cpErr != nil {
				s.logger.Errorf("err creating product:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", sixMonthPlan)
			}
		}
	}
}

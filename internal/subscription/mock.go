package subscription

import (
	"github.com/brianvoe/gofakeit"
	"math/rand"
	"time"
)

func (s service) DummyDataGenerator() error {
	s.userGenerator(5)
	s.productGenerator(5)
	s.planGenerator()
	s.voucherGenerator()
	s.voucherPlanGenerator()
	return nil
}

func (s service) userGenerator(count int) {
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

func (s service) productGenerator(count int) {
	for i := 0; i < count; i++ {
		name := gofakeit.DomainName()
		err := s.psql.CreateProduct(name)
		if err != nil {
			s.logger.Errorf("err creating product:%s", err)
		}
	}
}

func (s service) planGenerator() {
	products, err := s.GetProductsList()
	if err != nil {
		s.logger.Errorf("err fetching products:%s", err)
	} else {
		for i, product := range products {
			plan, cpErr := s.psql.CreatePlan("LifeTime", (i+1)*1000, 15, -1, product)
			if cpErr != nil {
				s.logger.Errorf("err creating plan:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", plan)
			}

			oneMonthPlan, cpErr := s.psql.CreatePlan("One Month", (i+1)*100, 2, 30, product)
			if cpErr != nil {
				s.logger.Errorf("err creating plan:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", oneMonthPlan)
			}

			threeMonthPlan, cpErr := s.psql.CreatePlan("Three Months", (i+1)*300, 6, 90, product)
			if cpErr != nil {
				s.logger.Errorf("err creating plan:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", threeMonthPlan)
			}

			sixMonthPlan, cpErr := s.psql.CreatePlan("Six Months", (i+1)*300, 8, 180, product)
			if cpErr != nil {
				s.logger.Errorf("err creating plan:%s", cpErr)
			} else {
				s.logger.Infof("PlanId:%d created", sixMonthPlan)
			}
		}
	}
}

func (s service) voucherGenerator() {
	plans, err := s.GetPlansList()
	if err != nil {
		s.logger.Errorf("err fetching products:%s", err)
	} else {
		discountTypes := []string{"amount", "percent"}
		for i, plan := range plans {
			name := gofakeit.BeerName()
			cvr := CreateVoucherRequest{
				Name:         name,
				PlanId:       plan.ID,
				Discount:     rand.Intn(30),
				DiscountType: discountTypes[i%2],
				StartDate:    time.Now(),
				EndDate:      time.Now().Add(time.Hour * 168),
			}
			_, err := s.psql.CreateVoucher(cvr)
			if err != nil {
				s.logger.Errorf("err creating voucher:%s", err)
			}

			//	make invalid voucher
			//name2 := gofakeit.BeerName()
			//cvr2 := CreateVoucherRequest{
			//	Name:         name2,
			//	PlanId:       plan.ID,
			//	Discount:     rand.Intn(30),
			//	DiscountType: discountTypes[i%2],
			//	StartDate:    time.Now().Add(-time.Hour * 100),
			//	EndDate:      time.Now().Add(-time.Hour * 24),
			//}
			//_, err2 := s.psql.CreateVoucher(cvr2)
			//if err2 != nil {
			//	s.logger.Errorf("err creating voucher:%s", err2)
			//}
		}
	}
}

func (s service) voucherPlanGenerator() {
	plans, err := s.psql.GetPlans()
	if err != nil {
		s.logger.Errorf("getting plans error : %s", err.Error())
		return
	}
	vouchers, err := s.psql.GetVouchers()
	if err != nil {
		s.logger.Errorf("getting vouchers error : %s", err.Error())
		return
	}
	rand.Seed(time.Now().Unix())
	//egt random list of numbers
	plansKeys := rand.Perm(len(plans))
	vouchersKeys := rand.Perm(len(vouchers))

	for i := 0; i < len(plans); i++ {
		pk := plansKeys[i]
		vk := vouchersKeys[i]
		_, err := s.psql.CreateVoucherPlan(plans[pk], vouchers[vk])
		if err != nil {
			s.logger.Errorf("creating voucher plans error : %s", err.Error())
			return
		}
	}
}

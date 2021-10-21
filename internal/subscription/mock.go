package subscription

import (
	"github.com/brianvoe/gofakeit"
	"math/rand"
	"sync"
	"time"
)

func (s service) DummyDataGenerator() error {
	wg := sync.WaitGroup{}
	wg.Add(5)
	go func() {
		defer wg.Done()
		s.userGenerator(10)
	}()
	go func() {
		defer wg.Done()
		s.productGenerator(10)
	}()
	go func() {
		defer wg.Done()
		s.planGenerator()
	}()
	go func() {
		defer wg.Done()
		s.voucherGenerator()
	}()
	go func() {
		defer wg.Done()
		s.voucherPlanGenerator(5)
	}()
	wg.Wait()
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

func (s service) voucherGenerator() {
	plans, err := s.GetPlansList()
	if err != nil {
		s.logger.Errorf("err fetching products:%s", err)
	} else {
		duscountTypes := []string{"amount", "percent"}
		for i, plan := range plans {
			name := gofakeit.BeerName()
			cvr := CreateVoucherRequest{
				Name:         name,
				PlanId:       plan.ID,
				Discount:     rand.Intn(30),
				DiscountType: duscountTypes[i%2],
				StartDate:    time.Now(),
				EndDate:      time.Now().Add(time.Duration(rand.Intn(200) * i)),
			}
			_, err := s.psql.CreateVoucher(cvr)
			if err != nil {
				s.logger.Errorf("err creating voucher:%s", err)
			}
		}
	}
}

func (s service) voucherPlanGenerator(count int) {
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
	plansKeys := rand.Perm(count)
	vouchersKeys := rand.Perm(count)

	for i := 0; i < count; i++ {
		_, err := s.psql.CreateVoucherPlan(plans[plansKeys[0]], vouchers[vouchersKeys[i]])
		if err != nil {
			s.logger.Errorf("creating voucher plans error : %s", err.Error())
			return
		}
	}
}

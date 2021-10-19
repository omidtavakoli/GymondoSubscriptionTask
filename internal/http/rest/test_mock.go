package rest

import (
	"Gymondo/internal/subscription"
	"github.com/stretchr/testify/mock"
)

type MockSubscriptionService struct {
	mock.Mock
}

func (m *MockSubscriptionService) GetProductById(id int) (subscription.Product, error) {
	args := m.Called(id)
	return args.Get(0).(subscription.Product), args.Error(1)
}

func (m *MockSubscriptionService) UserGenerator(count int) {
	_ = m.Called(count)
}
func (m *MockSubscriptionService) ProductGenerator(count int) {
	_ = m.Called(count)
}
func (m *MockSubscriptionService) PlanGenerator() {
	_ = m.Called()
}
func (m *MockSubscriptionService) GetProductsList() ([]subscription.Product, error) {
	args := m.Called()
	return args.Get(0).([]subscription.Product), args.Error(1)
}
func (m *MockSubscriptionService) BuyProduct(bpr subscription.BuyRequest) (subscription.UserPlan, error) {
	args := m.Called(bpr)
	return args.Get(0).(subscription.UserPlan), args.Error(1)
}
func (m *MockSubscriptionService) FetchPlansByUserId(userId int) ([]subscription.Status, error) {
	args := m.Called(userId)
	return args.Get(0).([]subscription.Status), args.Error(1)
}
func (m *MockSubscriptionService) ChangeUserPlanStatus(req subscription.ChangeStatus) error {
	args := m.Called(req)
	return args.Error(1)
}

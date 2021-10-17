package subscription

func (s service) GetProductsList() (products []Product, err error) {
	products, err = s.psql.GetProducts()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s service) GetProductById(id int) (product Product, err error) {
	product, err = s.psql.GetProduct(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s service) BuyProduct(bpr BuyRequest) (UserPlan, error) {
	plan, err := s.psql.BuyProduct(bpr)
	if err != nil {
		return plan, err
	}
	return plan, nil
}

func (s service) FetchPlansByUserId(userId int) ([]UserPlan, error){
	plans, err := s.psql.FetchPlansByUserId(userId)
	if err != nil {
		return plans, err
	}
	return plans, nil
}

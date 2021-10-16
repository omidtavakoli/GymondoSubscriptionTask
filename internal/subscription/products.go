package subscription

func (s service) GetProductsList() (products []Product, err error){
	products, err = s.psql.GetProducts()
	if err != nil {
		return products, err
	}
	return products,nil
}

package subscription

func (s service) ChangeUserPlanStatus(req ChangeStatus) error {
	return s.psql.ChangeUserPlanStatus(req)
}

func (s service) FetchPlansByUserId(userId int) ([]Status, error) {
	plans, err := s.psql.FetchPlansByUserId(userId)
	if err != nil {
		return plans, err
	}
	return plans, nil
}

func (s service) ProductByVoucher(voucherId int) ([]VoucherPlanProduct, error) {
	vpp, err := s.psql.FetchProductsByVoucherId(voucherId)
	if err != nil {
		return vpp, err
	}
	return vpp, nil
}

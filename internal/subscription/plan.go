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

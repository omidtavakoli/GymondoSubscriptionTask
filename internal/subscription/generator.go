package subscription

import "context"

func (s service) DataGenerator(ctx context.Context) {
	err := s.psql.GetUser("omdtvk@gmail.com")
	if err != nil {
		s.logger.Error(err)
	}
}

package user

import (
	"context"

)

type Service struct {
	storage Storage
	logger *logging.Logger
}

func (s *Service) Create(ctx context.Context, dta CreateUserDTO) (u User, err error) {
	//TODO next one
	return nil, nil
}
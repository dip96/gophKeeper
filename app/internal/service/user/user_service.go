package user

import (
	"context"
	"errors"
	auth "gophKeeper/internal/lib/auth/jwt"
	"gophKeeper/internal/uow"
)

type UserService struct {
	uow uow.UnitOfWork
}

func NewUserService(uow uow.UnitOfWork) *UserService {
	return &UserService{uow: uow}
}

func (s *UserService) RegisterUser(ctx context.Context, login, password string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	userRepo := s.uow.UserRepository()

	exists, err := userRepo.Exists(ctx, tx, login)
	if err != nil {
		return false, errors.New("failed to check user existence")
	}

	if exists {
		return false, errors.New("user already exists")
	}

	if err = userRepo.CreateUser(ctx, tx, login, password); err != nil {
		return false, errors.New("failed to create user")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *UserService) AuthenticateUser(ctx context.Context, login, password string) (string, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return "", errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	userRepo := s.uow.UserRepository()

	user, err := userRepo.GetUserByLogin(ctx, tx, login)
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	if err := s.uow.Commit(tx); err != nil {
		return "", errors.New("failed to commit transaction")
	}

	return token, nil
}

package login_data

import (
	"context"
	"errors"
	"gophKeeper/internal/models/entities"
	"gophKeeper/internal/uow"
)

type LoginDataService struct {
	uow uow.UnitOfWork
}

func NewLoginDataService(uow uow.UnitOfWork) *LoginDataService {
	return &LoginDataService{uow: uow}
}

func (s *LoginDataService) SaveLoginData(ctx context.Context, login, password string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	loginDataRepo := s.uow.LoginDataRepository()

	loginData := entities.LoginData{
		Username: login,
		Password: []byte(password),
	}

	if err = loginDataRepo.SaveData(ctx, tx, loginData); err != nil {
		return false, errors.New("failed to save login data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *LoginDataService) GetLoginData(ctx context.Context, entryID int) (*entities.LoginData, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	loginDataRepo := s.uow.LoginDataRepository()

	loginData, err := loginDataRepo.GetDataById(ctx, tx, entryID)
	if err != nil {
		return nil, err
	}

	if err := s.uow.Commit(tx); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &loginData, nil
}

func (s *LoginDataService) EditLoginData(ctx context.Context, entryID int, login, password string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	loginDataRepo := s.uow.LoginDataRepository()

	loginData := entities.LoginData{
		EntryID:  entryID,
		Username: login,
		Password: []byte(password),
	}

	if err = loginDataRepo.EditData(ctx, tx, entryID, loginData); err != nil {
		return false, errors.New("failed to edit login data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *LoginDataService) DeleteLoginData(ctx context.Context, entryID int) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	loginDataRepo := s.uow.LoginDataRepository()

	if err = loginDataRepo.DeleteData(ctx, tx, entryID); err != nil {
		return false, errors.New("failed to delete login data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

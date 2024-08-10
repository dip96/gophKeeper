package login_data

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/models/entities"
	"gophKeeper/internal/uow"
	pb "gophKeeper/protobuf/V1/login_data"
	"strconv"
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

	userID, ok := ctx.Value("user_id").(float64)
	if !ok {
		return false, errors.New("failed to save login data")
	}

	loginData := entities.LoginData{
		Username: login,
		Password: []byte(password),
		UserID:   int(userID),
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

func (s *LoginDataService) GetAllLoginData(ctx context.Context, page, limit int) ([]*pb.GetLoginDataResponse, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	loginDataRepo := s.uow.LoginDataRepository()
	allLoginData, err := loginDataRepo.GetAllData(ctx, tx, page, limit)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	var responseItems []*pb.GetLoginDataResponse
	for _, data := range allLoginData {
		responseItems = append(responseItems, &pb.GetLoginDataResponse{
			Login:    data.Username,
			Password: string(data.Password),
			Id:       strconv.Itoa(int(data.ID)),
		})
	}

	return responseItems, nil
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

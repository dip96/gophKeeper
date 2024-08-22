package credit_card_data

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/models/entities"
	"gophKeeper/internal/uow"
	pb "gophKeeper/protobuf/V1/credit_card_data"
	"strconv"
	"time"
)

type CreditCardDataService struct {
	uow uow.UnitOfWork
}

func NewCreditCardDataService(uow uow.UnitOfWork) *CreditCardDataService {
	return &CreditCardDataService{uow: uow}
}

func (s *CreditCardDataService) GetAllCreditCardData(ctx context.Context, page, limit int) ([]*pb.GetCreditCardDataResponse, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	dataRepo := s.uow.CreditCardDataRepository()
	allData, err := dataRepo.GetAllData(ctx, tx, page, limit)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	var responseItems []*pb.GetCreditCardDataResponse
	for _, data := range allData {
		expirationDate := data.ExpirationDate.Format("2006-01-02")
		responseItems = append(responseItems, &pb.GetCreditCardDataResponse{
			CardNumber:     data.CardNumber,
			CardholderName: data.CardholderName,
			ExpirationDate: expirationDate,
			Cvv:            data.CVV,
			Notes:          data.Notes,
			Id:             strconv.Itoa(data.ID),
		})
	}

	return responseItems, nil
}

func (s *CreditCardDataService) SaveCreditCardData(ctx context.Context, cardNumber, cardholderName, expirationDateStr, cvv, notes string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	userID, ok := ctx.Value("user_id").(float64)
	if !ok {
		return false, errors.New("failed to retrieve user id")
	}

	expirationDate, err := time.Parse("01/06", expirationDateStr)
	if err != nil {
		return false, status.Error(codes.InvalidArgument, "invalid expiration date format, expected YYYY-MM-DD")
	}

	creditCardDataRepo := s.uow.CreditCardDataRepository()

	creditCardData := entities.CreditCardData{
		CardNumber:     cardNumber,
		CardholderName: cardholderName,
		ExpirationDate: expirationDate,
		CVV:            cvv,
		Notes:          notes,
		UserID:         int(userID),
	}

	if err = creditCardDataRepo.SaveData(ctx, tx, creditCardData); err != nil {
		return false, errors.New("failed to save credit card data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *CreditCardDataService) GetCreditCardData(ctx context.Context, entryID int) (*entities.CreditCardData, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	creditCardDataRepo := s.uow.CreditCardDataRepository()

	creditCardData, err := creditCardDataRepo.GetDataById(ctx, tx, entryID)
	if err != nil {
		return nil, err
	}

	if err := s.uow.Commit(tx); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &creditCardData, nil
}

func (s *CreditCardDataService) EditCreditCardData(ctx context.Context, id int, cardNumber, cardholderName, expirationDateStr, cvv, notes string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	// Преобразуем строку в дату
	expirationDate, err := time.Parse("2006-01-02", expirationDateStr)
	if err != nil {
		return false, status.Error(codes.InvalidArgument, "invalid expiration date format, expected YYYY-MM-DD")
	}

	creditCardDataRepo := s.uow.CreditCardDataRepository()

	creditCardData := entities.CreditCardData{
		ID:             id,
		CardNumber:     cardNumber,
		CardholderName: cardholderName,
		ExpirationDate: expirationDate,
		CVV:            cvv,
		Notes:          notes,
	}

	if err = creditCardDataRepo.EditData(ctx, tx, id, creditCardData); err != nil {
		return false, errors.New("failed to edit credit card data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *CreditCardDataService) DeleteCreditCardData(ctx context.Context, entryID int) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	creditCardDataRepo := s.uow.CreditCardDataRepository()

	if err = creditCardDataRepo.DeleteData(ctx, tx, entryID); err != nil {
		return false, errors.New("failed to delete credit card data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

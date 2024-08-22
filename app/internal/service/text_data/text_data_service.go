package text_data

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/models/entities"
	"gophKeeper/internal/uow"
	pb "gophKeeper/protobuf/V1/text_data"
	"strconv"
)

type TextDataService struct {
	uow uow.UnitOfWork
}

func NewTextDataService(uow uow.UnitOfWork) *TextDataService {
	return &TextDataService{uow: uow}
}

func (s *TextDataService) GetAllTextData(ctx context.Context, page, limit int) ([]*pb.GetTextDataResponse, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	dataRepo := s.uow.TextDataRepository()
	allData, err := dataRepo.GetAllData(ctx, tx, page, limit)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	var responseItems []*pb.GetTextDataResponse
	for _, data := range allData {
		responseItems = append(responseItems, &pb.GetTextDataResponse{
			Text: data.Text,
			Id:   strconv.Itoa(data.ID),
		})
	}

	return responseItems, nil
}

func (s *TextDataService) SaveTextData(ctx context.Context, text string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	userID, ok := ctx.Value("user_id").(float64)
	fmt.Println(userID)
	if !ok {
		return false, errors.New("failed to save login data")
	}

	textDataRepo := s.uow.TextDataRepository()

	textData := entities.TextData{
		Text:   text,
		UserID: int(userID),
	}

	if err = textDataRepo.SaveData(ctx, tx, textData); err != nil {
		return false, errors.New("failed to save text data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *TextDataService) GetTextData(ctx context.Context, entryID int) (*entities.TextData, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	textDataRepo := s.uow.TextDataRepository()

	textData, err := textDataRepo.GetDataById(ctx, tx, entryID)
	if err != nil {
		return nil, err
	}

	if err := s.uow.Commit(tx); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &textData, nil
}

func (s *TextDataService) EditTextData(ctx context.Context, id int, text string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	textDataRepo := s.uow.TextDataRepository()

	textData := entities.TextData{
		ID:   id,
		Text: text,
	}

	if err = textDataRepo.EditData(ctx, tx, id, textData); err != nil {
		return false, errors.New("failed to edit text data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *TextDataService) DeleteTextData(ctx context.Context, entryID int) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	textDataRepo := s.uow.TextDataRepository()

	if err = textDataRepo.DeleteData(ctx, tx, entryID); err != nil {
		return false, errors.New("failed to delete text data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

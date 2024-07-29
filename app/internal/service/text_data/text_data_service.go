package text_data

import (
	"context"
	"errors"
	"gophKeeper/internal/models/entities"
	"gophKeeper/internal/uow"
)

type TextDataService struct {
	uow uow.UnitOfWork
}

func NewTextDataService(uow uow.UnitOfWork) *TextDataService {
	return &TextDataService{uow: uow}
}

func (s *TextDataService) SaveTextData(ctx context.Context, text string, entryID int) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	textDataRepo := s.uow.TextDataRepository()

	textData := entities.TextData{
		EntryID: entryID,
		Text:    text,
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

func (s *TextDataService) EditTextData(ctx context.Context, entryID int, text string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	textDataRepo := s.uow.TextDataRepository()

	textData := entities.TextData{
		EntryID: entryID,
		Text:    text,
	}

	if err = textDataRepo.EditData(ctx, tx, entryID, textData); err != nil {
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

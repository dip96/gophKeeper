package binary_data

import (
	"context"
	"errors"
	"gophKeeper/internal/models/entities"
	"gophKeeper/internal/uow"
)

type BinaryDataService struct {
	uow uow.UnitOfWork
}

func NewBinaryDataService(uow uow.UnitOfWork) *BinaryDataService {
	return &BinaryDataService{uow: uow}
}

func (s *BinaryDataService) SaveBinaryData(ctx context.Context, path, notes string, entryID int) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	binaryDataRepo := s.uow.BinaryDataRepository()

	binaryData := entities.BinaryData{
		EntryID: entryID,
		Path:    path,
		Notes:   notes,
	}

	if err = binaryDataRepo.SaveData(ctx, tx, binaryData); err != nil {
		return false, errors.New("failed to save binary data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *BinaryDataService) GetBinaryData(ctx context.Context, entryID int) (*entities.BinaryData, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return nil, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	binaryDataRepo := s.uow.BinaryDataRepository()

	binaryData, err := binaryDataRepo.GetDataById(ctx, tx, entryID)
	if err != nil {
		return nil, err
	}

	if err := s.uow.Commit(tx); err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &binaryData, nil
}

func (s *BinaryDataService) EditBinaryData(ctx context.Context, entryID int, path, notes string) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	binaryDataRepo := s.uow.BinaryDataRepository()

	binaryData := entities.BinaryData{
		EntryID: entryID,
		Path:    path,
		Notes:   notes,
	}

	if err = binaryDataRepo.EditData(ctx, tx, entryID, binaryData); err != nil {
		return false, errors.New("failed to edit binary data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

func (s *BinaryDataService) DeleteBinaryData(ctx context.Context, entryID int) (bool, error) {
	tx, err := s.uow.BeginTx(ctx)
	if err != nil {
		return false, errors.New("failed to begin transaction")
	}
	defer s.uow.Rollback(tx)

	binaryDataRepo := s.uow.BinaryDataRepository()

	if err = binaryDataRepo.DeleteData(ctx, tx, entryID); err != nil {
		return false, errors.New("failed to delete binary data")
	}

	if err := s.uow.Commit(tx); err != nil {
		return false, errors.New("failed to commit transaction")
	}

	return true, nil
}

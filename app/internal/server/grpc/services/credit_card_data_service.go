package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/service/credit_card_data"
	pb "gophKeeper/protobuf/V1/credit_card_data"
	"strconv"
)

type CreditCardDataService struct {
	pb.UnimplementedCreditCardDataServiceServer
	creditCardDataService *credit_card_data.CreditCardDataService
}

func NewCreditCardDataService(creditCardDataService *credit_card_data.CreditCardDataService) *CreditCardDataService {
	return &CreditCardDataService{creditCardDataService: creditCardDataService}
}

func (s *CreditCardDataService) GetAllCreditCardData(ctx context.Context, req *pb.GetAllCreditCardDataRequest) (*pb.GetAllCreditCardDataResponse, error) {
	responseItems, err := s.creditCardDataService.GetAllCreditCardData(ctx, 1, 10)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetAllCreditCardDataResponse{
		Items: responseItems,
	}, nil
}

func (s *CreditCardDataService) SaveCreditCardData(ctx context.Context, req *pb.CreditCardDataRequest) (*pb.CreditCardDataResponse, error) {
	if req.CardNumber == "" || req.CardholderName == "" || req.ExpirationDate == "" || req.Cvv == "" {
		return nil, status.Error(codes.InvalidArgument, "all fields are required")
	}

	success, err := s.creditCardDataService.SaveCreditCardData(ctx, req.CardNumber, req.CardholderName, req.ExpirationDate, req.Cvv, req.Notes)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.CreditCardDataResponse{Success: success}, nil
}

func (s *CreditCardDataService) GetCreditCardData(ctx context.Context, req *pb.GetCreditCardDataRequest) (*pb.GetCreditCardDataResponse, error) {
	entryID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	creditCardData, err := s.creditCardDataService.GetCreditCardData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetCreditCardDataResponse{
		CardNumber:     creditCardData.CardNumber,
		CardholderName: creditCardData.CardholderName,
		ExpirationDate: creditCardData.ExpirationDate.Format("01/06"),
		Cvv:            creditCardData.CVV,
		Notes:          creditCardData.Notes,
		Id:             strconv.Itoa(creditCardData.ID),
	}, nil
}

func (s *CreditCardDataService) EditCreditCardData(ctx context.Context, req *pb.EditCreditCardDataRequest) (*pb.CreditCardDataResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.creditCardDataService.EditCreditCardData(ctx, id, req.CardNumber, req.CardholderName, req.ExpirationDate, req.Cvv, req.Notes)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.CreditCardDataResponse{Success: success}, nil
}

func (s *CreditCardDataService) DeleteCreditCardData(ctx context.Context, req *pb.DeleteCreditCardDataRequest) (*pb.CreditCardDataResponse, error) {
	entryID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.creditCardDataService.DeleteCreditCardData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.CreditCardDataResponse{Success: success}, nil
}

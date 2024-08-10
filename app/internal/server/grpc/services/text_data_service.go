package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/service/text_data"
	pb "gophKeeper/protobuf/V1/text_data"
	"strconv"
)

type TextDataService struct {
	pb.UnimplementedTextDataServiceServer
	textDataService *text_data.TextDataService
}

func NewTextDataService(textDataService *text_data.TextDataService) *TextDataService {
	return &TextDataService{textDataService: textDataService}
}

func (s *TextDataService) GetAllTextData(ctx context.Context, req *pb.GetAllTextDataRequest) (*pb.GetAllTextDataResponse, error) {
	responseItems, err := s.textDataService.GetAllTextData(ctx, 1, 10)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetAllTextDataResponse{
		Items: responseItems,
	}, nil
}

func (s *TextDataService) SaveTextData(ctx context.Context, req *pb.TextDataRequest) (*pb.TextDataResponse, error) {
	if req.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "text is required")
	}

	success, err := s.textDataService.SaveTextData(ctx, req.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.TextDataResponse{Success: success}, nil
}

func (s *TextDataService) GetTextData(ctx context.Context, req *pb.GetTextDataRequest) (*pb.GetTextDataResponse, error) {
	entryID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	textData, err := s.textDataService.GetTextData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetTextDataResponse{
		Text: textData.Text,
		Id:   strconv.Itoa(textData.ID),
	}, nil
}

func (s *TextDataService) EditTextData(ctx context.Context, req *pb.EditTextDataRequest) (*pb.TextDataResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.textDataService.EditTextData(ctx, id, req.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.TextDataResponse{Success: success}, nil
}

func (s *TextDataService) DeleteTextData(ctx context.Context, req *pb.DeleteTextDataRequest) (*pb.TextDataResponse, error) {
	entryID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.textDataService.DeleteTextData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.TextDataResponse{Success: success}, nil
}

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

func (s *TextDataService) SaveTextData(ctx context.Context, req *pb.TextDataRequest) (*pb.TextDataResponse, error) {
	if req.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "text is required")
	}

	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.textDataService.SaveTextData(ctx, req.Text, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.TextDataResponse{Success: success}, nil
}

func (s *TextDataService) GetTextData(ctx context.Context, req *pb.GetTextDataRequest) (*pb.GetTextDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	textData, err := s.textDataService.GetTextData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetTextDataResponse{
		Text:    textData.Text,
		EntryId: strconv.Itoa(textData.EntryID),
	}, nil
}

func (s *TextDataService) EditTextData(ctx context.Context, req *pb.EditTextDataRequest) (*pb.TextDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.textDataService.EditTextData(ctx, entryID, req.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.TextDataResponse{Success: success}, nil
}

func (s *TextDataService) DeleteTextData(ctx context.Context, req *pb.DeleteTextDataRequest) (*pb.TextDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.textDataService.DeleteTextData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.TextDataResponse{Success: success}, nil
}

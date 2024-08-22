package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/service/binary_data"
	pb "gophKeeper/protobuf/V1/binary_data"
	"strconv"
)

type BinaryDataService struct {
	pb.UnimplementedBinaryDataServiceServer
	binaryDataService *binary_data.BinaryDataService
}

func NewBinaryDataService(binaryDataService *binary_data.BinaryDataService) *BinaryDataService {
	return &BinaryDataService{binaryDataService: binaryDataService}
}

func (s *BinaryDataService) SaveBinaryData(ctx context.Context, req *pb.BinaryDataRequest) (*pb.BinaryDataResponse, error) {
	if req.Path == "" {
		return nil, status.Error(codes.InvalidArgument, "path is required")
	}

	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.binaryDataService.SaveBinaryData(ctx, req.Path, req.Notes, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.BinaryDataResponse{Success: success}, nil
}

func (s *BinaryDataService) GetBinaryData(ctx context.Context, req *pb.GetBinaryDataRequest) (*pb.GetBinaryDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	binaryData, err := s.binaryDataService.GetBinaryData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetBinaryDataResponse{
		Path:    binaryData.Path,
		Notes:   binaryData.Notes,
		EntryId: strconv.Itoa(binaryData.EntryID),
	}, nil
}

func (s *BinaryDataService) EditBinaryData(ctx context.Context, req *pb.EditBinaryDataRequest) (*pb.BinaryDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.binaryDataService.EditBinaryData(ctx, entryID, req.Path, req.Notes)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.BinaryDataResponse{Success: success}, nil
}

func (s *BinaryDataService) DeleteBinaryData(ctx context.Context, req *pb.DeleteBinaryDataRequest) (*pb.BinaryDataResponse, error) {
	entryID, err := strconv.Atoi(req.EntryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid entry_id")
	}

	success, err := s.binaryDataService.DeleteBinaryData(ctx, entryID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.BinaryDataResponse{Success: success}, nil
}

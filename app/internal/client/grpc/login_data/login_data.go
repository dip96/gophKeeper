package login_data

import (
	"context"
	"gophKeeper/internal/client/grpc"
	"gophKeeper/protobuf/V1/login_data"
	pb "gophKeeper/protobuf/V1/login_data"
)

func GetAllLoginData() (*pb.GetAllLoginDataResponse, error) {
	conn, err := grpc.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &login_data.GetAllLoginDataRequest{}

	items, err := client.GetAllLoginData(context.Background(), req)
	return items, err
}

func SaveLoginData(login, password, entryID string) (*pb.LoginDataResponse, error) {
	conn, err := grpc.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &login_data.LoginDataRequest{
		Login:    login,
		Password: password,
		EntryId:  entryID,
	}

	resp, err := client.SaveLoginData(context.Background(), req)
	return resp, err
}

func GetLoginData(entryID string) (*pb.GetLoginDataResponse, error) {
	conn, err := grpc.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &login_data.GetLoginDataRequest{
		EntryId: entryID,
	}

	resp, err := client.GetLoginData(context.Background(), req)
	return resp, err
}

func EditLoginData(login, password, entryID string) (*pb.LoginDataResponse, error) {
	conn, err := grpc.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &login_data.EditLoginDataRequest{
		Login:    login,
		Password: password,
		EntryId:  entryID,
	}

	resp, err := client.EditLoginData(context.Background(), req)
	return resp, err
}

func DeleteLoginData(entryID string) (*pb.LoginDataResponse, error) {
	conn, err := grpc.GetConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &login_data.DeleteLoginDataRequest{
		EntryId: entryID,
	}

	resp, err := client.DeleteLoginData(context.Background(), req)
	return resp, err
}

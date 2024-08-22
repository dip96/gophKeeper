package login_data

import (
	"context"
	"fmt"
	"gophKeeper/internal/client/grpc"
	pb "gophKeeper/protobuf/V1/login_data"
)

type LoginDataRepository struct{}

type LoginDataRequest struct {
	Login    string
	Password string
	Id       string
}

func (r *LoginDataRepository) GetAllData() ([]LoginDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &pb.GetAllLoginDataRequest{}
	resp, err := client.GetAllLoginData(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var result []LoginDataRequest
	for _, item := range resp.Items {
		result = append(result, LoginDataRequest{
			Login:    item.Login,
			Password: item.Password,
			Id:       item.Id,
		})
	}
	return result, nil
}

func (r *LoginDataRepository) SaveData(data LoginDataRequest) (LoginDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return LoginDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &pb.LoginDataRequest{
		Login:    data.Login,
		Password: data.Password,
		EntryId:  data.Id,
	}

	resp, err := client.SaveLoginData(context.Background(), req)
	if err != nil {
		return LoginDataRequest{}, err
	}

	if !resp.Success {
		return LoginDataRequest{}, fmt.Errorf("failed to save data")
	}

	return data, nil
}

func (r *LoginDataRepository) GetData(entryID string) (LoginDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return LoginDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &pb.GetLoginDataRequest{
		EntryId: entryID,
	}

	resp, err := client.GetLoginData(context.Background(), req)
	if err != nil {
		return LoginDataRequest{}, err
	}

	return LoginDataRequest{
		Login:    resp.Login,
		Password: resp.Password,
		Id:       resp.EntryId,
	}, nil
}

func (r *LoginDataRepository) EditData(data LoginDataRequest) (LoginDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return LoginDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &pb.EditLoginDataRequest{
		Login:    data.Login,
		Password: data.Password,
		EntryId:  data.Id,
		Id:       data.Id,
	}

	resp, err := client.EditLoginData(context.Background(), req)
	if err != nil {
		return LoginDataRequest{}, err
	}

	if !resp.Success {
		return LoginDataRequest{}, fmt.Errorf("failed to edit data")
	}

	return data, nil
}

func (r *LoginDataRepository) DeleteData(entryID string) (LoginDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return LoginDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewLoginDataServiceClient(conn)

	req := &pb.DeleteLoginDataRequest{
		EntryId: entryID,
	}

	resp, err := client.DeleteLoginData(context.Background(), req)
	if err != nil {
		return LoginDataRequest{}, err
	}

	if !resp.Success {
		return LoginDataRequest{}, fmt.Errorf("failed to delete data")
	}

	return LoginDataRequest{Id: entryID}, nil
}

func (r LoginDataRequest) GetID() string {
	return r.Id
}

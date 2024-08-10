package text_data

import (
	"context"
	"gophKeeper/internal/client/grpc"
	pb "gophKeeper/protobuf/V1/text_data"
)

type TextDataRequest struct {
	Text string
	Id   string
}

type TextDataRepository struct{}

func (r *TextDataRepository) GetAllData() ([]TextDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewTextDataServiceClient(conn)

	req := &pb.GetAllTextDataRequest{}
	resp, err := client.GetAllTextData(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var result []TextDataRequest
	for _, item := range resp.Items {
		result = append(result, TextDataRequest{
			Text: item.Text,
			Id:   item.Id,
		})
	}
	return result, nil
}

func (r *TextDataRepository) SaveData(data TextDataRequest) (TextDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return TextDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewTextDataServiceClient(conn)

	req := &pb.TextDataRequest{
		Text: data.Text,
		Id:   data.Id,
	}

	_, err = client.SaveTextData(context.Background(), req)
	if err != nil {
		return TextDataRequest{}, err
	}

	return data, nil
}

func (r *TextDataRepository) GetData(entryID string) (TextDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return TextDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewTextDataServiceClient(conn)

	req := &pb.GetTextDataRequest{
		Id: entryID,
	}

	resp, err := client.GetTextData(context.Background(), req)
	if err != nil {
		return TextDataRequest{}, err
	}

	return TextDataRequest{
		Text: resp.Text,
		Id:   resp.Id,
	}, nil
}

func (r *TextDataRepository) EditData(data TextDataRequest) (TextDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return TextDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewTextDataServiceClient(conn)

	req := &pb.EditTextDataRequest{
		Text: data.Text,
		Id:   data.Id,
	}

	_, err = client.EditTextData(context.Background(), req)
	if err != nil {
		return TextDataRequest{}, err
	}

	return data, nil
}

func (r *TextDataRepository) DeleteData(entryID string) (TextDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return TextDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewTextDataServiceClient(conn)

	req := &pb.DeleteTextDataRequest{
		Id: entryID,
	}

	_, err = client.DeleteTextData(context.Background(), req)
	if err != nil {
		return TextDataRequest{}, err
	}

	return TextDataRequest{Id: entryID}, nil
}

func (r TextDataRequest) GetID() string {
	return r.Id
}

package credit_card_data

import (
	"context"
	"gophKeeper/internal/client/grpc"
	pb "gophKeeper/protobuf/V1/credit_card_data"
)

type CreditCardDataRequest struct {
	CardholderName string
	CardNumber     string
	ExpirationDate string
	CVV            string
	Notes          string
	Id             string
}

type CreditCardDataRepository struct{}

func (r *CreditCardDataRepository) GetAllData() ([]CreditCardDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewCreditCardDataServiceClient(conn)

	req := &pb.GetAllCreditCardDataRequest{}
	resp, err := client.GetAllCreditCardData(context.Background(), req)
	if err != nil {
		return nil, err
	}

	var result []CreditCardDataRequest
	for _, item := range resp.Items {
		result = append(result, CreditCardDataRequest{
			CardholderName: item.CardholderName,
			CardNumber:     item.CardNumber,
			ExpirationDate: item.ExpirationDate,
			CVV:            item.Cvv,
			Notes:          item.Notes,
			Id:             item.Id,
		})
	}
	return result, nil
}

func (r *CreditCardDataRepository) SaveData(data CreditCardDataRequest) (CreditCardDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return CreditCardDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewCreditCardDataServiceClient(conn)

	req := &pb.CreditCardDataRequest{
		CardholderName: data.CardholderName,
		CardNumber:     data.CardNumber,
		ExpirationDate: data.ExpirationDate,
		Cvv:            data.CVV,
		Notes:          data.Notes,
		Id:             data.Id,
	}

	_, err = client.SaveCreditCardData(context.Background(), req)
	if err != nil {
		return CreditCardDataRequest{}, err
	}

	return data, nil
}

func (r *CreditCardDataRepository) GetData(entryID string) (CreditCardDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return CreditCardDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewCreditCardDataServiceClient(conn)

	req := &pb.GetCreditCardDataRequest{
		Id: entryID,
	}

	resp, err := client.GetCreditCardData(context.Background(), req)
	if err != nil {
		return CreditCardDataRequest{}, err
	}

	return CreditCardDataRequest{
		CardholderName: resp.CardholderName,
		CardNumber:     resp.CardNumber,
		ExpirationDate: resp.ExpirationDate,
		CVV:            resp.Cvv,
		Notes:          resp.Notes,
		Id:             resp.Id,
	}, nil
}

func (r *CreditCardDataRepository) EditData(data CreditCardDataRequest) (CreditCardDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return CreditCardDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewCreditCardDataServiceClient(conn)

	req := &pb.EditCreditCardDataRequest{
		CardholderName: data.CardholderName,
		CardNumber:     data.CardNumber,
		ExpirationDate: data.ExpirationDate,
		Cvv:            data.CVV,
		Notes:          data.Notes,
		Id:             data.Id,
	}

	_, err = client.EditCreditCardData(context.Background(), req)
	if err != nil {
		return CreditCardDataRequest{}, err
	}

	return data, nil
}

func (r *CreditCardDataRepository) DeleteData(entryID string) (CreditCardDataRequest, error) {
	conn, err := grpc.GetConnAuthenticated()
	if err != nil {
		return CreditCardDataRequest{}, err
	}
	defer conn.Close()
	client := pb.NewCreditCardDataServiceClient(conn)

	req := &pb.DeleteCreditCardDataRequest{
		Id: entryID,
	}

	_, err = client.DeleteCreditCardData(context.Background(), req)
	if err != nil {
		return CreditCardDataRequest{}, err
	}

	return CreditCardDataRequest{Id: entryID}, nil
}

func (r CreditCardDataRequest) GetID() string {
	return r.Id
}

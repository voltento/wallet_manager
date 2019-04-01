package payment

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/database"
)
import "context"

func MakeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(changeBalanceRequest)
		v, err := svc.changeBalance(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeSendMoneyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(sendMoneyRequest)
		v, err := svc.sendMoney(req)
		if err != nil {
			return database.Error{Msg: "Error occured during exec SendMoney query", Error: err.Error()}, nil
		}
		return v, nil
	}
}
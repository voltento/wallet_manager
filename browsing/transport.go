package browsing

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/pursesManager/database"
	"net/http"
)

func DecodeRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type accResponse = database.Account
type paymentResponse = database.Payment

func MakeGetAccountsHandler(s Service) http.Handler {
	return kithttp.NewServer(
		MakeGetAccountsEndpoint(s),
		DecodeRequest,
		EncodeResponse,
	)
}

func MakeGetPaymentsHandler(s Service) http.Handler {
	return kithttp.NewServer(
		MakeGetPaymentsEndpoint(s),
		DecodeRequest,
		EncodeResponse,
	)
}
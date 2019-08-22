package transport

import (
	"cargo-handling/datastruct"
	dt "cargo-handling/datastruct"
	"cargo-handling/service"
	"context"
	"encoding/json"
	"log"

	//"kit/endpoint"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type HandlingService interface {
	GetHandleProcessService(context.Context, dt.Handle) []dt.Handle
}

type handlingService struct{}

func (handlingService) GetHandleProcessService(_ context.Context, a dt.Handle) []dt.Handle {
	return call_GetHandleProcessService(a)
}

func call_GetHandleProcessService(a dt.Handle) []dt.Handle {

	messageResponse := service.GetHandlingProcess(a)

	return messageResponse
}

func makeGetHandleProcessEndpoint(handling HandlingService) endpoint.Endpoint {
	log.Println("Process")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.GetHandleProcessRequest)
		paramDel := dt.Handle{}
		paramDel.ROUTE_ID = req.ROUTE_ID
		paramDel.ITENARY_ID = req.ITENARY_ID
		paramDel.ROUTING_STATUS = req.ROUTING_STATUS
		handling.GetHandleProcessService(ctx, paramDel)
		return datastruct.GetHandleProcessResponse{
			// ROUTE_ID:       1,
			// ITENARY_ID:     1,
			ROUTING_STATUS: "UNLOAD",
		}, nil
	}
}

func decodeGetHandleProcessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.GetHandleProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	handling := handlingService{}

	GetHandleProcessHandler := httptransport.NewServer(
		makeGetHandleProcessEndpoint(handling),
		decodeGetHandleProcessRequest,
		encodeResponse,
	)
	//url path of our API service
	http.Handle("/GetHandleProcess", GetHandleProcessHandler)

}

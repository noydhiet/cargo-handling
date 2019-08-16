package transport

import (
	"cargo-handling/datastruct"
	dt "cargo-handling/datastruct"
	svc "cargo-handling/service"
	"context"
	"encoding/json"
	"log"

	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	UpdateStatusHandlingService(context.Context, dt.Handle) []dt.Handle
}

type aphService struct {}

func (aphService) UpdateStatusHandlingService(_ context.Context, del dt.Handle) []dt.Handle {
	return call_ServiceUpdateStatusHandlingService(del)
}

func call_ServiceUpdateStatusHandlingService(del dt.Handle) []dt.Handle {
	retDel := svc.UpdateStatusHandling(del)

	return retDel
}

func makeUpdateStatusHandlingEndpoint(aph AphService) endpoint.Endpoint {
	log.Println("Process")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.UpdateStatusHandlingRequest)
		paramDel := dt.Handle{}
		paramDel.ID_ROUTE = req.ROUTE_ID
		paramDel.ID_ITENARY = req.ITENARY_ID
		paramDel.ROUTING_STATUS = req.ROUTING_STATUS
		paramDel.TRANSPORT_STATUS = req.TRANSPORT_STATUS
		aph.UpdateStatusHandlingService(ctx, paramDel)
		return datastruct.UpdateStatusHandlingResponse{
			ROUTING_STATUS: "Success",
			TRANSPORT_STATUS: "Success",
		}, nil
	}
}

func decodeUpdateHandlingStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.UpdateStatusHandlingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServiceAndStartListener() {
	aph := aphService{}

	UpdateStatusHandlingHandler := httptransport.NewServer(
		makeUpdateStatusHandlingEndpoint(aph),
		decodeUpdateHandlingStatusRequest,
		encodeResponse,
	)

	//url path of our API Service
	http.Handle("UpdateStatusHandling", UpdateStatusHandlingHandler)
}
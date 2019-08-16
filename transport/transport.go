package transport

import (
	"cargo-Handling/datastruct"
	dt "cargo-Handling/datastruct"
	svc "cargo-Handling/service"
	"context"
	"encoding/json"
	"log"

	//"kit/endpoint"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	UpdateStatusHandlingService(context.Context, dt.Handling) []dt.Handling
}

type aphService struct{

}

func (aphService) UpdateStatusHandlingService(_ context.Context, del dt.Handling) []dt.Handling {
	return call_UpdateStatusHandlingService(hand)
}

func call_ServiceUpdateStatusHandlingService(hand dt.Handling) []dt.Handling {
	retHand := svc.HandStatusHandling(hand)

	return retHand
}

func makeUpdateStatusHandlingEndpoint(aph AphService) endpoint.Endpoint {
	log.Println("Process")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.UpdateStatusHandlingRequest)
		paramHand := dt.Handling{}
		paramHand.ID_ROUTE = req.ROUTE_ID
		paramHand.ID_ITENARY = req.ITENARY_ID
		paramHand.STATUS_outing_status = req.ROUTING_STATUS
		paramHand.STATUS_TRANSPORT = req.TRANSPORT_STATUS
		aph.UpdateStatusHandlingService(ctx, paramHand)
		return datastruct.GetStatusDeliveryResponse{
			ROUTING_STATUS:      "Success",
			TRANSPORT_STATUS:    "Success",
		}, nil
	}
}

func decodeUpdateHandlingStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.GetStatusDeliveryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	GetStatusDeliveryHandler := httptransport.NewServer(
		makeGetStatusDeliveryEndpoint(aph),
		decodeGetDeliveryStatusRequest,
		encodeResponse,
	)
	//url path of our API service
	http.Handle("/GetStatusDelivery", GetStatusDeliveryHandler)

}
package datastruct

type UpdateStatusHandlingRequest struct {
	ROUTE_ID int `json:"id_route_spec"`
	ITENARY_ID int `json:"id_itenary"`
	ROUTING_STATUS string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
}

type UpdateStatusHandlingResponse struct {
	ROUTING_STATUS string `json:"routing_status"`
}

type GetHandleProcessRequest struct {
	ROUTE_ID int `json:"id_route_spec"`
	ITENARY_ID int `json:"id_itenary"`
	VOYAGE_NUMBER int `json:"voyage_number"`
}

type GetHandleProcessResponse struct {
	ROUTING_STATUS string `json:"routing_status"`
}

type Handle struct {
	ID_ROUTE int 
	ID_ITENARY int 
	ROUTING_STATUS string 
	TRANSPORT_STATUS string 
}



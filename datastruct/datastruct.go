package datastruct

type UpdateStatusHandlingRequest struct {
	ROUTE_ID         int    `json:"id_route_spec"`
	ITENARY_ID       int    `json:"id_itenary"`
	ROUTING_STATUS   string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
}

type UpdateStatusHandlingResponse struct {
	STATUS  int    `json:"status"`
	MESSAGE string `json:"message"`
}

type GetStatusHandlingRequest struct {
	ROUTE_ID      string `json:"id_route_spec"`
	ITENARY_ID    string `json:"id_itenary"`
	VOYAGE_NUMBER string `json:"voyage_number"`
}

type GetStatusHandlingResponse struct {
	ROUTING_STATUS string `json:"routing_status"`
}

type Handling struct {
	ID_ROUTE      string
	ID_ITENARY    string
	NUMBER_VOYAGE string
}

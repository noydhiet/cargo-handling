package datastruct

type UpdateStatusHandlingRequest struct {
	ROUTE_ID         int    `json:"id_route_spec"`
	ITENARY_ID       int    `json:"id_itenary"`
	ROUTING_STATUS   string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
}

type UpdateStatusHandlingResponse struct {
	STATUS  string `json:"status"`
	MESSAGE string `json:"message"`
}

type GetHandleProcessRequest struct {
	ROUTE_ID      int `json:"id_route_spec"`
	ITENARY_ID    int `json:"id_itenary"`
	VOYAGE_NUMBER int `json:"voyage_number"`
}

type GetHandleProcessResponse struct {
	ROUTING_STATUS string `json:"routing_status"`
}

type Handle struct {
	ROUTING_STATUS      string
	TRANSPORT_STATUS    string
	LAST_KNOWN_LOCATION string
	ID_ITENARY          int
	ID_ROUTE            int
}

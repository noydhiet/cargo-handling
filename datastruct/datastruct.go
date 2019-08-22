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
	ROUTE_ID       int    `json:"fk_id_route_spec"`
	ITENARY_ID     int    `json:"fk_id_itenary"`
	ROUTING_STATUS string `json:"routing_status"`
}

type GetHandleProcessResponse struct {
	ROUTING_STATUS string `json:"routing_status"`
	// ROUTE_ID            int    `json:"fk_id_route_spec"`
	// ITENARY_ID          int    `json:"fk_id_itenary"`
	// TRANSPORT_STATUS    string `json:"transport_status"`
	// DELIVERY_ID         int    `json:"idt_delivery"`
	// LAST_KNOWN_LOCATION string `json:"last_known_location"`
}

type Handle struct {
	ROUTE_ID            int
	ITENARY_ID          int
	ROUTING_STATUS      string
	TRANSPORT_STATUS    string
	DELIVERY_ID         int
	LAST_KNOWN_LOCATION string
}

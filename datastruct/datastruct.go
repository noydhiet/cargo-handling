package datastruct

type UpdateStatusHandlingRequest struct {
	ROUTE_ID         int    `json:"id_route_spec"`
	ITENARY_ID       int    `json:"id_itenary"`
	ROUTING_STATUS   string `json:"routing_status"`
	TRANSPORT_STATUS string `json:"transport_status"`
}

type UpdateStatusHandlingResponse struct {
<<<<<<< HEAD
	STATUS  int    `json:"status"`
=======
	STATUS string `json:"status"`
>>>>>>> 5456f15dacda1d4d83b1720fa248eb88e8fce12f
	MESSAGE string `json:"message"`
}

type GetStatusHandlingRequest struct {
<<<<<<< HEAD
	ROUTE_ID      string `json:"id_route_spec"`
	ITENARY_ID    string `json:"id_itenary"`
	VOYAGE_NUMBER string `json:"voyage_number"`
=======
	ROUTE_ID int `json:"id_route_spec"`
	ITENARY_ID int `json:"id_itenary"`
	VOYAGE_NUMBER int `json:"voyage_number"`
>>>>>>> 5456f15dacda1d4d83b1720fa248eb88e8fce12f
}

type GetStatusHandlingResponse struct {
	ROUTING_STATUS string `json:"routing_status"`
}

<<<<<<< HEAD
type Handling struct {
	ID_ROUTE      string
	ID_ITENARY    string
	NUMBER_VOYAGE string
=======
type Handling struct{
	ROUTING_STATUS      string
	TRANSPORT_STATUS    string
	LAST_KNOWN_LOCATION string
	ID_ITENARY          int
	ID_ROUTE            int
>>>>>>> 5456f15dacda1d4d83b1720fa248eb88e8fce12f
}

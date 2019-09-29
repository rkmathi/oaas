package httprequestresponse

type GetBalanceResponse struct {
	Amount int32 `json:"amount"`
}
type UpdateBalanceResponse struct {
	Status string `json:"status"`
	Amount int32  `json:"amount"`
}

type CreateOmikujiResponse struct {
	Status string `json:"status"`
	Value  string `json:"value"`
}

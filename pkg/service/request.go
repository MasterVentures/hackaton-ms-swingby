package service

//CreateTransactionRequest defines fields to create transaction
type CreateReqRequest struct {
}

//CreateTransactionResponse defines the request response
type CreateReqResponse struct {
	ID  string `json:"id"`
	Err string `json:"err"`
}

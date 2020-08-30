package service

//CreateTransactionRequest defines fields to create transaction
type CreateDepositRequest struct {
}

//CreateTransactionResponse defines the request response
type CreateDepositResponse struct {
	ID  string `json:"id"`
	Err string `json:"err"`
}

package service

//CreateTransactionRequest defines fields to create transaction
type CreateWithdrawalRequest struct {
	Addresses []Addresses `json:"addresses"`
}

type Addresses struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

//CreateTransactionResponse defines the request response
type CreateWithdrawalResponse struct {
	ID  string `json:"id"`
	Err string `json:"err"`
}

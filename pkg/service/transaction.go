package service

//CreateTransactionRequest defines fields to create transaction
type CreateTransactionRequest struct {
	Input         []Input  `json:"input"`
	Output        []Output `json:"output"`
	Confirmations int      `json:"confirmations"`
}

type Input struct {
	Address string `json:"address"`
}

type Output struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

//CreateTransactionResponse defines the request response
type CreateTransactionResponse struct {
	ID  string `json:"id"`
	Err string `json:"err"`
}

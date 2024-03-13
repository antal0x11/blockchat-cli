package dst

type Transaction struct {
	RecipientAddress  string  `json:"recipient_address"`
	TypeOfTransaction string  `json:"type_of_transaction"`
	Amount            float64 `json:"amount,omitempty"`
	Message           string  `json:"message,omitempty"`
}

type BalanceResponse struct {
	Status    string  `json:"status"`
	Reason    string  `json:"reason,omitempty"`
	Timestamp string  `json:"timestamp"`
	Balance   float64 `json:"balance,omitempty"`
}

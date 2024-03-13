package dst

type Transaction struct {
	SenderAddress     string  `json:"sender_address,omitempty"`
	RecipientAddress  string  `json:"recipient_address"`
	TypeOfTransaction string  `json:"type_of_transaction"`
	Amount            float64 `json:"amount,omitempty"`
	Message           string  `json:"message,omitempty"`
	Fee               float64 `json:"fee,omitempty"`
	Nonce             uint32  `json:"nonce,omitempty"`
	TransactionId     string  `json:"transaction_id,omitempty"`
	Signature         []byte  `json:"signature,omitempty"`
}

type BalanceResponse struct {
	Status    string  `json:"status"`
	Reason    string  `json:"reason,omitempty"`
	Timestamp string  `json:"timestamp"`
	Balance   float64 `json:"balance,omitempty"`
}

type Block struct {
	Index        uint32        `json:"index,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Validator    string        `json:"validator,omitempty"`
	Hash         string        `json:"hash,omitempty"`
	PreviousHash string        `json:"previous_hash,omitempty"`
	Capacity     uint32        `json:"capacity,omitempty"`
}

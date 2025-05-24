package model

type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prevHash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
	Difficulty   int           `json:"difficulty"`
}

type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

type Wallet struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

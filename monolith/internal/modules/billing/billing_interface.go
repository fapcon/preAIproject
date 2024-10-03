package billing

type Billinger interface {
	GetBalances(in GetBalancesIn) GetWalletOut
	GetCryptoBalance() BOut
	TopUpBalance(in TopUpBalanceIn) BOut
	TopDownBalance(in TopDownBalanceIn) BOut
	GetOperationsHistory(in GetOperationsHistoryIn) GetOperationsHistoryOut
	Transfer(in TransferIn) TransferOut
}

type TransferIn struct {
	UserID         int
	FromCurrency   int
	ToCurrency     int
	Source         int
	Type           int
	Description    string
	Amount         float64
	ExpectedAmount float64
}

type TransferOut struct {
	UserID         int
	OriginalAmount float64
	Transferred    float64
	ErrorCode      int
}

type GetOperationsHistoryIn struct {
	UserID   int
	DateFrom string
	DateTo   string
	Limit    int
}

type GetOperationsHistoryOut struct {
	UserID     int
	Operations []Operation
	ErrorCode  int
}

type Operation struct {
	Amount      float64
	Type        int
	Source      int
	Description string
	Date        string
}

type TopUpBalanceIn struct {
	UserID      int
	Amount      float64
	Type        int
	Source      int
	Description string
}

type TopDownBalanceIn struct {
	UserID      int
	Amount      float64
	Type        int
	Source      int
	Description string
}

type GetBalancesIn struct {
	UserID int
}

type Wallet struct {
	Crypto map[int]float64
	Fiat   map[int]float64
}

type GetWalletOut struct {
	Wallet    Wallet
	ErrorCode int
}

type BOut struct {
	ErrorCode int
}

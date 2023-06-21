package dto

type TradeInput struct {
	OrderID       string  `json:"order_id"`
	InvestorID    string  `json:"investor_id"`
	AssetID       string  `json:"asset_id"`
	CurrentShares int64   `json:"current_shares"`
	Shares        int64   `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"order_type"`
}

type OrderOutput struct {
	OrderID           string               `json:"order_id"`
	InvestorID        string               `json:"investor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         string               `json:"order_type"`
	Status            string               `json:"status"`
	Partial           int64                 `json:"partial"`
	Shares            int64                `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerID       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int64   `json:"shares"`
}

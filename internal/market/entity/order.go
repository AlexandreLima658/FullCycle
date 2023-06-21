package entity

type Order struct {
	ID            string         `json:"id"`
	Investor      *Investor      `json:"investor"`
	Asset         *Asset         `json:"asset"`
	Shares        int64          `json:"shares"`
	PendingShares int64          `json:"pending_shares"`
	Price         float64        `json:"price"`
	OrderType     string         `json:"order_type"`
	Status        string         `json:"status"`
	Transactions  []*Transaction `json:"transections"`
}

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int64, price float64, orderType string) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}

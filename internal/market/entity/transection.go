package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string    `json:"id"`
	SellingOrder *Order    `json:"selling_order"`
	BuyingOrder  *Order    `json:"buying_order"`
	Shares       int64     `json:"shares"`
	Price        float64   `json:"price"`
	Total        float64   `json:"total"`
	DateTime     time.Time `json:"date_time"`
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int64, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int64, price float64) {
	t.Total = float64(t.Shares) * t.BuyingOrder.Price
}

func (t *Transaction) CloseBuyOrder() {
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
} 

func (t *Transaction) CloseSellOrder() {
	if t.BuyingOrder.PendingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) AddBuyOrderPendingShares(shares int64) {
	t.BuyingOrder.PendingShares += shares

}
func (t *Transaction) AddSellOrderPendingShares(shares int64) {
	t.SellingOrder.PendingShares += shares
}

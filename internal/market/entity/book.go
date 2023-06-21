package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order         []*Order        `json:"orders"`
	Transections  []*Transaction  `json:"transections"`
	OrdersChan    chan *Order     `json:"orders_chan"`
	OrdersChanOut chan *Order     `json:"orders_chan_out"`
	Wg            *sync.WaitGroup `json:"wg"`
}

func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:         []*Order{},
		Transections:  []*Transaction{},
		OrdersChan:    orderChan,
		OrdersChanOut: orderChanOut,
		Wg:            wg,
	}
}

func (b *Book) Trade() {

	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	for order := range b.OrdersChan {
		asset := order.Asset.ID

		if buyOrders[asset] == nil {
			buyOrders[asset] = NewOrderQueue()
			heap.Init(buyOrders[asset])
		}
		if sellOrders[asset] == nil {
			sellOrders[asset] = NewOrderQueue()
			heap.Init(sellOrders[asset])
		}
		if order.OrderType == "BUY" {
			buyOrders[asset].Push(order)
			if sellOrders[asset].Len() > 0 && sellOrders[asset].Orders[0].Price <= order.Price {
				sellOrder := sellOrders[asset].Pop().(*Order)
				if sellOrder.PendingShares > 0 {
					transection := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					b.AddTransection(transection, b.Wg)
					sellOrder.Transactions = append(sellOrder.Transactions, transection)
					order.Transactions = append(order.Transactions, transection)
					b.OrdersChanOut <- sellOrder
					b.OrdersChanOut <- order
					if sellOrder.PendingShares > 0 {
						sellOrders[asset].Push(sellOrder)
					}
				}
			} else if order.OrderType == "SELL" {
				sellOrders[asset].Push(order)
				if buyOrders[asset].Len() > 0 && buyOrders[asset].Orders[0].Price >= order.Price {
					buyOrder := buyOrders[asset].Pop().(*Order)
					if buyOrder.PendingShares > 0 {
						transection := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)
						b.AddTransection(transection, b.Wg)
						buyOrder.Transactions = append(buyOrder.Transactions, transection)
						order.Transactions = append(order.Transactions, transection)
						b.OrdersChanOut <- buyOrder
						b.OrdersChanOut <- order
						if buyOrder.PendingShares > 0 {
							buyOrders[asset].Push(buyOrder)
						}
					}

				}
			}

		}
	}

}

func (b *Book) AddTransection(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	sellingShares := transaction.SellingOrder.PendingShares
	buyingShares := transaction.BuyingOrder.PendingShares

	minShares := sellingShares
	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.ID, -minShares)
	transaction.AddSellOrderPendingShares(-minShares)
	transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.ID, minShares)
	transaction.AddBuyOrderPendingShares(-minShares)

	transaction.CalculateTotal(transaction.Shares, transaction.Price)

	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()

	b.Transections = append(b.Transections, transaction)
}

package orderbook

import (
	"fmt"
	"sort"
)

type Orderbook struct {
	Ask []*Order
	Bid []*Order
}

func New() *Orderbook {
	return &Orderbook{}
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	return orderbook.Trade(order)
}

func (ob *Orderbook) Trade(order *Order) ([]*Trade, *Order) {
	tradeArr := []*Trade{}
	var propArr *[]*Order

	switch order.Side {
	case SideAsk:
		propArr = &ob.Bid
	case SideBid:
		propArr = &ob.Ask
	}

	for order.Volume != 0 {
		if len(*propArr) == 0 {
			ob.RejectOrder(order)
			break
		}

		theBestProp := (*propArr)[0]

		if propAccepted(order, theBestProp, 0) {
			var tradeValue uint64
			if order.Volume < theBestProp.Volume {
				tradeValue = order.Volume
			} else {
				tradeValue = theBestProp.Volume
			}

			trade := Trade{theBestProp, order, tradeValue, theBestProp.Price}
			order.Volume -= tradeValue
			theBestProp.Volume -= tradeValue
			tradeArr = append(tradeArr, &trade)

			if theBestProp.Volume == 0 {
				*propArr = (*propArr)[1:]
			}
		} else {
			ob.RejectOrder(order)
			break
		}
	}

	if order.Volume > 0 && order.Kind == KindMarket {
		return tradeArr, order
	}

	return tradeArr, nil
}

func propAccepted(order *Order, prop *Order, side int) bool {
	switch order.Kind {
	case KindMarket:
		return true
	case KindLimit:
		switch order.Side {
		case SideAsk:
			return order.Price <= prop.Price
		case SideBid:
			return order.Price >= prop.Price
		}
	}

	fmt.Println("Inaccessible point was accessed")
	return false
}

func (ob *Orderbook) RejectOrder(order *Order) {
	switch order.Side {
	case SideAsk:
		ob.Ask = append(ob.Ask, order)
		sort.Slice(ob.Ask, func(i, j int) bool {
			return ob.Ask[i].Price < ob.Ask[j].Price
		})
	case SideBid:
		ob.Bid = append(ob.Bid, order)
		sort.Slice(ob.Bid, func(i, j int) bool {
			return ob.Bid[i].Price > ob.Bid[j].Price
		})
	}
}

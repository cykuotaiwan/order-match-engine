package engine

import (
	"fmt"
)

type Order struct {
	OrderId uint64
	Dir     Direction
	Price   uint64
	Amount  uint64
}

func (order *Order) Print() {
	fmt.Printf("%d, %d, %d, %d\n", order.OrderId, order.Dir, order.Price, order.Amount)
}

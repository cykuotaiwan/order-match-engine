package main

import (
	"fmt"
	"order-match-engine/engine"
)

func main() {
	fmt.Println("starter")
	order := engine.Order{OrderId: 20, Dir: engine.SELL, Price: 200, Amount: 10}
	order.Print()
}

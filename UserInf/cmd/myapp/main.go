package main

import (
	"fmt"
	"myapp/internal"
)

var price int

func main() {
	fmt.Printf("Введите стоимость покупки: ")
	fmt.Scan(&price)

	cust := internal.NewCustomer("Dmitry", 23, 10000, 100, false)
	discount, err := cust.CalcDiscount()
	cost, _ := internal.CalcPrice(*cust, price)
	fmt.Println(discount, err)
	fmt.Println(cost)
}

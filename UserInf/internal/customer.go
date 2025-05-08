package internal

import (
	"errors"
)

type Customer struct {
	Name             string
	Age              int
	Balance          int
	Debt             int
	Discount         bool
	CalcDiscount     func() (int, error)
	CalcPrice        func(Customer, int) (int, error)
	DefaultDiscount int
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	c := &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
		DefaultDiscount: 500,
	}
	c.CalcDiscount = func() (int, error) {
		if !c.Discount {
			return 0, errors.New("discount not available")
		}
		result := c.DefaultDiscount - c.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	c.CalcPrice = CalcPrice
	return c
}

func CalcPrice(c Customer, price int) (int, error) {
	discount, err := c.CalcDiscount()
	if err != nil {
		price = 0
		return price, err
	}
	return price - discount, err
}

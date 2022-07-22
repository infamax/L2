package main

import (
	"errors"
	"fmt"
)

// Паттерн цепочка обязанностей
// Реализуем на примере выборы платежной системы

type Handler interface {
	CanPay(amount float64) error
}

// MasterCard Account
type MasterCard struct {
	balance float64
	next    Handler
}

func NewMasterCard(balance float64, next Handler) *MasterCard {
	return &MasterCard{
		balance: balance,
		next:    next,
	}
}

func (m *MasterCard) CanPay(amount float64) error {
	if m.balance > amount {
		m.balance -= amount
		fmt.Println("Successfully payment using mastercard")
		return nil
	} else if m.next != nil {
		return m.next.CanPay(amount)
	} else {
		return errors.New("all accounts have been checked and there are not enough funds on each")
	}
}

type Qiwi struct {
	balance float64
	next    Handler
}

func NewQiwiPayment(balance float64, next Handler) *Qiwi {
	return &Qiwi{
		balance: balance,
		next:    next,
	}
}

func (q *Qiwi) CanPay(amount float64) error {
	if q.balance > amount {
		q.balance -= amount
		fmt.Println("successfully payment using qiwi")
		return nil
	} else if q.next != nil {
		return q.next.CanPay(amount)
	} else {
		return errors.New("all accounts have been checked and there are not enough funds on each")
	}
}

type PayPal struct {
	balance float64
	next    Handler
}

func NewPayPal(balance float64, next Handler) *PayPal {
	return &PayPal{
		balance: balance,
		next:    next,
	}
}

func (p *PayPal) CanPay(amount float64) error {
	if p.balance > amount {
		p.balance -= amount
		fmt.Println("successfully payment using paypal payment")
		return nil
	} else if p.next != nil {
		return p.next.CanPay(amount)
	} else {
		return errors.New("all accounts have been checked and there are not enough funds on each")
	}
}

func main() {
	masterCard := NewMasterCard(120, nil)
	qiwiPayment := NewQiwiPayment(55, masterCard)
	paypal := NewPayPal(23, qiwiPayment)
	fmt.Println(paypal.CanPay(100))
}

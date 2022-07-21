package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// Реализовать паттерн фасад

// Система доставка

// Courier Курьер
type Courier struct {
	billingID int
	name      string
	position  string
	// some information has here, about male, age, height and etc
}

func NewCourier(billingID int, name string) *Courier {
	return &Courier{
		billingID: billingID,
		name:      name,
	}
}

func (c *Courier) Deliver(address string) {
	log.Printf("Order has successfully delivered by address: %s\n", address)
}

func (c *Courier) BillingID() int {
	return c.billingID
}

func (c *Courier) Position() string {
	return c.position
}

func (c *Courier) ChangeName(newName string) {
	log.Printf("Courier change name: %s on new name: %s\n", c.name, newName)
	c.name = newName
}

// Recipient Получатель заказа
type Recipient struct {
	billingID int
	name      string
	city      string
	address   string
}

func NewRecipient(billingID int, name, city, address string) *Recipient {
	return &Recipient{
		billingID: billingID,
		name:      name,
		city:      city,
		address:   address,
	}
}

func (r *Recipient) BillingID() int {
	return r.billingID
}

func (r *Recipient) Address() string {
	return r.address
}

func (r *Recipient) Order(price float64, address string) {
	log.Println("Recipient has order")
}

func (r *Recipient) ChangeLocation(city, address string) {
	log.Println("Recipient change location on new location")
	r.city = city
	r.address = address
}

// Billing Счет
type Billing struct {
	balance float64
}

func NewBilling(balance float64) *Billing {
	return &Billing{balance: balance}
}

func (b *Billing) Balance() float64 {
	return b.balance
}

func (b *Billing) Deposit(amount float64) {
	b.balance += amount
}

func (b *Billing) Withdraw(amount float64) error {
	if amount > b.balance {
		return errors.New("not enough money")
	}
	b.balance -= amount
	return nil
}

// DeliverySystem система доставки
type DeliverySystem struct {
	couriers   []Courier
	recipients []Recipient
	billings   []Billing
}

func NewDeliverySystem() *DeliverySystem {
	return &DeliverySystem{
		couriers:   []Courier{},
		recipients: []Recipient{},
		billings:   []Billing{},
	}
}

func (ds *DeliverySystem) NewRecipient(name, city, address string) int {
	billing := NewBilling(0)
	ds.billings = append(ds.billings, *billing)
	recipient := NewRecipient(len(ds.billings)-1, name, city, address)
	ds.recipients = append(ds.recipients, *recipient)
	return len(ds.recipients) - 1
}

func (ds *DeliverySystem) NewCourier(name string) {
	billing := NewBilling(0)
	ds.billings = append(ds.billings, *billing)
	courier := NewCourier(len(ds.billings)-1, name)
	ds.couriers = append(ds.couriers, *courier)
}

func (ds *DeliverySystem) NewOrder(clientID int) {
	// Находим свободного курьера и ближайшего к адресу
	// для простоты реализации просто берем id нулевого курьера
	cost := calculateCost(ds.couriers[0].Position(), ds.recipients[clientID].Address())
	// Проверяем что у клиента достаточно средств для совершения заказа
	if err := ds.billings[ds.recipients[clientID].BillingID()].Withdraw(cost); err != nil {
		log.Println("not enough money")
	}
	// Выполняем заказ
	ds.recipients[clientID].Order(cost, ds.recipients[clientID].Address())
	ds.couriers[0].Deliver(ds.recipients[clientID].Address())
	// Начисляем деньги курьеру за заказ
	ds.billings[ds.couriers[0].BillingID()].Deposit(cost)
}

func calculateCost(src, dst string) float64 {
	return math.Abs(float64(len(src))-float64(len(dst))) * 15
}

func main() {
	ds := NewDeliverySystem()
	ds.NewCourier("vasya")
	id := ds.NewRecipient("petya", "unknown", "unknown")
	fmt.Println("id = ", id)
	ds.NewOrder(id)
}

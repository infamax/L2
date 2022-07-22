package main

import (
	"fmt"
	"math"
)

// Паттерн стратегия

// Будем реализовать разную систему доставки
// Доставку пешком, на личном авто, на грузовом авто

type DeliveryCoastI interface {
	CalculateCoast(src, dst string) float64 // Расчет стоимости доставки
	// Для каждого типа доставки свой тип расчета стоимости доставки
}

type WalkedDelivery struct {
	coef float64
}

func NewWalkedDelivery(coef float64) *WalkedDelivery {
	return &WalkedDelivery{
		coef: coef,
	}
}

func (wd *WalkedDelivery) CalculateCoast(src, dst string) float64 {
	if len(src) == len(dst) {
		return wd.coef
	}
	return math.Abs(float64(len(src)-len(dst))) * wd.coef
}

type AutoDelivery struct {
	coef float64
}

func NewAutoDelivery(coef float64) *AutoDelivery {
	return &AutoDelivery{coef: coef}
}

func (ad *AutoDelivery) CalculateCoast(src, dst string) float64 {
	return float64(len(src)+len(dst)) * ad.coef
}

type CargoDelivery struct {
	coef float64
}

func NewCargoDelivery(coef float64) *CargoDelivery {
	return &CargoDelivery{coef: coef}
}

func (cd *CargoDelivery) CalculateCoast(src, dst string) float64 {
	return float64(len(src)*len(dst)) * cd.coef
}

func main() {
	var deliveries []DeliveryCoastI
	for i := 0; i < 5; i++ {
		if i%3 == 0 {
			deliveries = append(deliveries, NewWalkedDelivery(1.1))
		} else if i%3 == 1 {
			deliveries = append(deliveries, NewAutoDelivery(1.3))
		} else {
			deliveries = append(deliveries, NewCargoDelivery(2.2))
		}
	}

	fmt.Println("Walk = ", deliveries[0].CalculateCoast("abc", "dfg"))
	fmt.Println("Auto = ", deliveries[1].CalculateCoast("abc", "dfg"))
	fmt.Println("Cargo = ", deliveries[2].CalculateCoast("abc", "dfg"))
}

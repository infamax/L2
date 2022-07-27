package objects

import "fmt"

type Garage struct {
	address string
	open    bool
}

func NewGarage(address string) *Garage {
	return &Garage{
		address: address,
		open:    false,
	}
}

func (g *Garage) OpenDoor() {
	fmt.Printf("Garage open door in this address: %s\n", g.address)
	g.open = true
}

func (g *Garage) Close() {
	fmt.Printf("Garage close door in this adrress: %s\n", g.address)
	g.open = false
}

package commands

import "github.com/infamax/l2/patterns/command/objects"

type GarageCloseDoorCommand struct {
	garage *objects.Garage
}

func NewGarageCloseDoorCommand(garage *objects.Garage) *GarageCloseDoorCommand {
	return &GarageCloseDoorCommand{
		garage: garage,
	}
}

func (g *GarageCloseDoorCommand) Execute() {
	g.garage.Close()
}

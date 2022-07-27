package commands

import "github.com/infamax/l2/patterns/command/objects"

type GarageOpenDoorCommand struct {
	garage *objects.Garage
}

func NewGarageOpenDoorCommand(garage *objects.Garage) *GarageOpenDoorCommand {
	return &GarageOpenDoorCommand{
		garage: garage,
	}
}

func (garageOpen *GarageOpenDoorCommand) Execute() {
	garageOpen.garage.OpenDoor()
}

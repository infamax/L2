package commands

import "github.com/infamax/l2/patterns/command/objects"

type LightOnCommand struct {
	light *objects.Light
}

func NewLightOnCommand(light *objects.Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (lon *LightOnCommand) Execute() {
	lon.light.On()
}

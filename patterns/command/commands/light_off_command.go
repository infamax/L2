package commands

import "github.com/infamax/l2/patterns/command/objects"

type LightOffCommand struct {
	light *objects.Light
}

func NewLightOffCommand(light *objects.Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (lof *LightOffCommand) Execute() {
	lof.light.Off()
}

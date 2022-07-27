package remote_control

import (
	"fmt"
	"github.com/infamax/l2/patterns/command/commands"
)

type RemoteControl struct {
	commandsOn  []commands.Command
	commandsOff []commands.Command
}

func NewRemoteControl(countCommand int) *RemoteControl {
	return &RemoteControl{
		commandsOn:  make([]commands.Command, countCommand, countCommand),
		commandsOff: make([]commands.Command, countCommand, countCommand),
	}
}

func (r *RemoteControl) SetCommandOn(numberButton int, command commands.Command) {
	fmt.Printf("Set on command on button: %d\n", numberButton)
	r.commandsOn[numberButton-1] = command
}

func (r *RemoteControl) SetCommandOff(numberButton int, command commands.Command) {
	fmt.Printf("Set off command on button: %d\n", numberButton)
	r.commandsOff[numberButton-1] = command
}

func (r *RemoteControl) PressedButtonOn(numberButton int) {
	fmt.Printf("Pressed Onbutton number: %d\n", numberButton)
	r.commandsOn[numberButton-1].Execute()
}

func (r *RemoteControl) PressedButtonOff(numberButton int) {
	fmt.Printf("Pressed Offbutton number: %d\n", numberButton)
	r.commandsOff[numberButton-1].Execute()
}

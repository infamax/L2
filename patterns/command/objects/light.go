package objects

import "fmt"

type Light struct {
	room  string
	model string
	on    bool
}

func NewLight(room, model string) *Light {
	return &Light{
		room:  room,
		model: model,
		on:    false,
	}
}

func (l *Light) On() {
	fmt.Printf("Light on in the %s\n", l.room)
	l.on = true
}

func (l *Light) Off() {
	fmt.Printf("Light off in the %s\n", l.room)
	l.on = false
}

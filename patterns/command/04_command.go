package main

import (
	"github.com/infamax/l2/patterns/command/commands"
	"github.com/infamax/l2/patterns/command/objects"
	"github.com/infamax/l2/patterns/command/remote_control"
)

// Реализация паттерна команда
// На примере системы умного дома
// Система управления умного дома может поддерживать различные команды
// Такие как включения и выключения света, открытие и закрытия двери гаража,
// и тому подобное. Чтобы в будущем не было проблем с добавлением новых команда,
// как раз используем паттерн команда

func main() {
	remoteControl := remote_control.NewRemoteControl(2)   // Пульт управляет командами, но ничего не знает об их содержимом
	light := objects.NewLight("kitchen", "energy saving") // Создание объекта лампочки
	lightOn := commands.NewLightOnCommand(light)          // Команда включения лампочки
	lightOff := commands.NewLightOffCommand(light)        // Команда выключения лампочки
	remoteControl.SetCommandOn(1, lightOn)                // Установка команды на пульте
	remoteControl.SetCommandOff(1, lightOff)

	garage := objects.NewGarage("somewhere")                  // объект гаража
	garageOpen := commands.NewGarageOpenDoorCommand(garage)   // команда открытия гаража
	garageClose := commands.NewGarageCloseDoorCommand(garage) // команда закрытия гаража
	remoteControl.SetCommandOn(2, garageOpen)                 // установка соответствующих команд на пульте
	remoteControl.SetCommandOff(2, garageClose)

	// Моделируем пример работы
	// Пользователь вышел из дома на работу
	// Ему нужно выключить свет и открыть дверь гаража, чтобы забрать машину и затем закрыть ее

	remoteControl.PressedButtonOff(1)
	remoteControl.PressedButtonOn(2)
	remoteControl.PressedButtonOff(2)

	// Теперь пользователь вернулся из работы, соответственно
	// ему нужно нажать кнопку вернуть машину в гараж и включить свет

	remoteControl.PressedButtonOn(2)
	remoteControl.PressedButtonOff(2)
	remoteControl.PressedButtonOn(1)
}

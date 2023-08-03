package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		fmt.Println("Down")
		time.Sleep(time.Millisecond * 500)

		led.High()
		fmt.Println("Up")
		time.Sleep(time.Millisecond * 500)
	}
}

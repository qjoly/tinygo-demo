package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	led := machine.GPIO26
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		led.High()
		fmt.Println("LED ON")
		time.Sleep(1000 * time.Millisecond)

		led.Low()
		fmt.Println("LED OFF")
		time.Sleep(500 * time.Millisecond)
	}
}

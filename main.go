package main

import (
	"fmt"

	"github.com/yonkolevel/ableton-push-sdk/push2"
)

func main() {
	c := make(chan int)
	push := push2.New()
	push.OpenSession()
	push.Display.Open()
	push.SendCC(uint8(push2.Play), 126)
	push.SendCC(uint8(push2.Record), 126)
	push.SendCC(uint8(push2.Automate), 126)
	push.SendCC(uint8(push2.FixedLength), 126)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Print("Hello world")
}

// TODO: leds, fix having to coerce controlids uint8(push2.Play), read all midi message types, screen

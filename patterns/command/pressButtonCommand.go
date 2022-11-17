package main

type PressPowerButtonCommand struct {
	device Device
}

func (c *PressPowerButtonCommand) execute() {
	c.device.powerOn()
}

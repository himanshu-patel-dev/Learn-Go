package main

// Interface

type Command interface {
	execute()
}

// Concrete Implementation - On command

type OnCommand struct {
	device Device
}

func (o *OnCommand) execute() {
	o.device.on()
}

// Concrete Implementation - Off command

type OffCommand struct {
	device Device
}

func (f *OffCommand) execute() {
	f.device.off()
}

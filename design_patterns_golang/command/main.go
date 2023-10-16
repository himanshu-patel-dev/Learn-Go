package main

func main() {
	tv := &TV{}

	onCommand := OnCommand{device: tv}
	offCommand := OffCommand{device: tv}

	// Button expect pointer receiver while
	// calling press() method on it, but even if
	// we use value receiver, golang by default find
	// the pointer to value and use that to call the
	// pointer receiver method onButton.press()
	// will get pointer to Button (&Button) and press()
	// can be called on pointer of Button

	// But while passing onCommand or offCommand
	// we pass pointer to them, showhow the value of
	// these are not converted to pointer and thus we get error

	// cannot use offCommand (variable of type OffCommand) as
	// type Command in struct literal:
	// OffCommand does not implement Command (execute method
	// has pointer receiver)

	onButton := Button{command: &onCommand}
	onButton.press()

	offButton := Button{command: &offCommand}
	offButton.press()
}

//TV is On
//TV is Off

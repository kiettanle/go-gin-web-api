package main

func main() {
	tv := &Tv{isRunning: false}

	// onCommand := &OnCommand{
	// 	device: tv,
	// }

	// offCommand := &OffCommand{
	// 	device: tv,
	// }

	// onButton := &Button{
	// 	cmd: onCommand,
	// }
	// onButton.press()

	// offButton := &Button{
	// 	cmd: offCommand,
	// }

	// offButton.press()

	pressPowerButtonCommand := &PressPowerButtonCommand{device: tv}

	powerButton := &Button{cmd: pressPowerButtonCommand}

	powerButton.press()
	powerButton.press()
	powerButton.press()
	powerButton.press()
	powerButton.press()
	powerButton.press()
}

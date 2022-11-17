package main

type Button struct {
	cmd Command
}

func (b *Button) press() {
	b.cmd.execute()
}

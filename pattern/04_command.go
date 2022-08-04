package pattern

import "fmt"

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

type Command interface {
	Execute()
}

type PlayCommand struct {
	device1 device1
}

func (c *PlayCommand) Execute() {
	c.device1.play()
}

type PauseCommand struct {
	device1 device1
}

func (c *PauseCommand) Execute() {
	c.device1.pause()
}

type device1 interface {
	play()
	pause()
}

type Walkman struct {
	songIsPlaying bool
}

func (w *Walkman) play() {
	w.songIsPlaying = true
	fmt.Println("Song is playing...")
}
func (w *Walkman) pause() {
	w.songIsPlaying = false
	fmt.Println("Song on pause...")
}

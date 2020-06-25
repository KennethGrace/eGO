package engine

import (
	"time"
)

const (
	TargetFps int = 60
	TargetUps int = 30
)

type engine struct {
	fps int
	ups int
	win *Window
	timer *time.Timer
	game Logic
}

func NewEngine(title string, width, height int, vsync bool, game Logic) *engine {
	e := &engine{fps: TargetFps, ups: TargetUps, game: game}
	e.win = &Window{title: title, width: width, height: height, vsync: vsync}
	return e
}

func (e engine) Run() {
	if err := e.init(); err != nil {
		panic(err)
	}
	e.loop()
	e.cleanup()
}

func (e engine) init() error {
	if err := e.win.init(); err != nil {
		return err
	}
	return nil
}

func (e engine) loop() {
	var elapsedTime float64
	var accumulator float64 = 0
	interval := float64(1 / e.ups)
	running := true
	for running && !e.win.shouldClose() {
		elapsedTime = getElapsedTime()
		accumulator += elapsedTime
		e.input()
		for ;accumulator >= interval; accumulator -= interval {
			e.update(interval)
		}
		e.render()
		if !e.win.vsync {
			e.sync()
		}
	}
}

func (e engine) sync() {

}

func (e engine) input() {

}

func (e engine) update(interval float64) {

}

func (e engine) render() {

}

func (e engine) cleanup() {

}

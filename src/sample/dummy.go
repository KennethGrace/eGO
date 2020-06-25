package sample

import "eGO/src/engine"

type game struct {
	renderer *engine.Renderer
}

func NewGame() *game {
	return &game{&engine.Renderer{}}
}

func (g game) Init(engine.Window) error {
	if err := g.renderer.Init(); err != nil {
		return err
	}
	return nil
}

func (g game) Input(window engine.Window) {

}

func (g game) Update(interval float32) {

}

func (g game) Render(window engine.Window) {

}

func (g game) Cleanup() {

}

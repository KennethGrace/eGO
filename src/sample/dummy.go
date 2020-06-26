package sample

import "eGO/src/engine"

type game struct {
	renderer *engine.Renderer
}

func NewGame() *game {
	return &game{&engine.Renderer{}}
}

func (g game) Init(window engine.Window) error {
	if err := g.renderer.Init(window); err != nil {
		return err
	}
	return nil
}

func (g game) Input(window engine.Window) {

}

func (g game) Update(interval float64) {

}

func (g game) Render(window engine.Window) {

}

func (g game) Cleanup() {

}

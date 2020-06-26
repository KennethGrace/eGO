package engine

type Renderer struct {
	w Window
}

func (r Renderer) Init(window Window) error {
	return nil
}

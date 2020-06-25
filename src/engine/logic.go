package engine

type Logic interface {
	Init(Window) error
	Input(Window)
	Update(float32)
	Render(Window)
	Cleanup()
}

package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var lastLoopTime float64

func getTime() float64 {
	return glfw.GetTime() / 1000000000
}

func getElapsedTime() float64 {
	time := getTime()
	elapsedTime := time - lastLoopTime
	lastLoopTime = time
	return elapsedTime
}

func getLastLoopTime() float64 {
	return lastLoopTime
}

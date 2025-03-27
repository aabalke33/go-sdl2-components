package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	FPS = 60.0
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
        "SDL Component System",
        sdl.WINDOWPOS_UNDEFINED,
        sdl.WINDOWPOS_UNDEFINED,
        800,
        600,
        sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	w, h := window.GetSize()

	scene := NewScene(renderer, FPS, w, h, 10, sdl.Color{
		R: 25, G: 25, B: 25, A: 255},
	)

	//scene.Add(NewBox(scene, 0, 0, 200, 200, 3, sdl.Color{
	//	R: 100, G: 50, B: 100, A: 255},
	//))

    ratio := 1/3.0
	scene.Add(NewBoxPercentage(scene, ratio, &scene.H, 0, 0, 2, sdl.Color{
		R: 194, G: 138, B: 51, A: 255},
	))
	scene.Add(NewTexturedBox(renderer, scene, 1, &scene.H, 0, 0, 1, 5, 5))

	for scene.Active {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		scene.Update(1/FPS, nil)
		scene.View(renderer)
		renderer.Present()
	}
}

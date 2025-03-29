package main

import (
	"time"

	components "github.com/aabalke33/go-sdl2-components/Components"
	"github.com/aabalke33/go-sdl2-components/examples"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	FPS = 60.0
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()
	defer sdl.Quit()
	defer ttf.Quit()

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

	scene := components.NewScene(renderer, w, h, 10, sdl.Color{
		R: 25, G: 25, B: 25, A: 255},
	)
	text := examples.NewText(renderer, scene, 200, 200, 0, 0, 5, "Created by Aaron Balke", 24)
	scene.Add(text)

	perm := examples.NewBox(scene, 0, 0, 200, 200, 3, sdl.Color{
		R: 100, G: 50, B: 100, A: 255},
	)
	text = examples.NewText(renderer, perm, 200, 200, 0, 0, 7, "2", 48)
	perm.Add(text)
	scene.Add(perm)

	ratio := 1 / 3.0
	perc := examples.NewBoxPercentage(scene, ratio, &scene.H, 0, 0, 2, sdl.Color{
		R: 194, G: 138, B: 51, A: 255},
	)
	text = examples.NewText(renderer, perc, 200, 200, 0, 0, 7, "3", 12)
	perc.Add(text)
	scene.Add(perc)

	//scene.Add(components.NewTexturedBox(renderer, scene, 1, &scene.H, 0, 0, 1, 5, 5))

	frameTime := time.Second / 60 //FPS
	ticker := time.NewTicker(frameTime)
    for i := range ticker.C {

		if !scene.Active {
			ticker.Stop()
			break
		}

        if i.UnixMicro() % 7 == 0 {
            scene.DeleteInactive()
            //println(i.UnixMicro())
        }

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		scene.Update(1/FPS, nil)
		scene.View(renderer)
		renderer.Present()


	}
}

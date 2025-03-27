package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Box struct {
	parent        *Component
	x, y, w, h, z int32
	color         sdl.Color
	Active        bool
}

func NewBox(parent Component, x, y, w, h, z int32, color sdl.Color) *Box {
	return &Box{
		parent: &parent,
		x:      x,
		y:      y,
		w:      w,
		h:      h,
		z:      z,
		color:  color,
		Active: true,
	}
}

func (b *Box) Update(dt float64, event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseButtonEvent:
		if e.Type == sdl.MOUSEBUTTONDOWN {
			b.Active = !b.Active
		}
	}
}

func (b *Box) View(renderer *sdl.Renderer) {
	if !b.Active {
		return
	}

	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	rect := sdl.Rect{X: b.x, Y: b.y, W: b.w, H: b.h}
	renderer.FillRect(&rect)
}

func (b *Box) GetZ() int32 {
	return b.z
}

func (b *Box) Resize() {
	return
}

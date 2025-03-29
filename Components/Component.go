package components

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Component interface {
	Update(dt float64, event sdl.Event)
	View(renderer *sdl.Renderer)
	Add(component Component)

	Resize()

	GetZ() int32
	IsActive() bool

	GetChildren() []*Component
    SetChildren([]*Component)
	GetParent() *Component

	GetSize() (h int32, w int32)
}

package components

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Box struct {
	parent        *Component
    children      []*Component
	x, y, w, h, z int32
	color         sdl.Color
	Active        bool
}

func NewBox(parent Component, x, y, w, h, z int32, color sdl.Color) *Box {
	return &Box{
		parent: &parent,
        children: []*Component{},
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


    if !(*b.parent).isActive() {
        b.Active = false
    }
}

func (b *Box) View(renderer *sdl.Renderer) {
	if !b.Active {
		return
	}

	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	rect := sdl.Rect{X: b.x, Y: b.y, W: b.w, H: b.h}
	renderer.FillRect(&rect)

	countRendered := 0
	var z int32 = 0
	for z = range 100 {
		if len(b.children) == countRendered {
			return
		}

		for _, c := range b.children {
            cp := *c
			if cp.GetZ() == z {
				cp.View(renderer)
			}
		}
	}





}

func (b *Box) Add(c Component) {
	b.children = append(b.children, &c)
}

func (b *Box) GetZ() int32 {
	return b.z
}

func (b *Box) Resize() {
	return
}

func (b *Box) isActive() bool {
    return b.Active
}

func (b *Box) GetChildren() []*Component {
    return b.children
}

func (b *Box) GetParent() *Component {
    return b.parent
}

func (b *Box) GetSize() (int32, int32) {
    return b.h, b.w
}

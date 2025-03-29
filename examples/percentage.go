package examples

import (
	"github.com/veandco/go-sdl2/sdl"
	comp "github.com/aabalke33/go-sdl2-components/Components"
	"math"
)

type BoxPercentage struct {
	children   []*comp.Component
	parent     *comp.Component
	w, x, y, z int32
	h          *int32
	ratio      float64
	color      sdl.Color
	Active     bool
}

func NewBoxPercentage(parent comp.Component, ratio float64, h *int32, x, y, z int32, color sdl.Color) *BoxPercentage {

	b := BoxPercentage{
		parent: &parent,
		ratio:  ratio,
		x:      x,
		y:      y,
		h:      h,
		z:      z,
		color:  color,
		Active: true,
	}

	b.Resize()

	return &b
}

func (b *BoxPercentage) Update(dt float64, event sdl.Event) {

	if !b.Active {
		return
	}

	switch e := event.(type) {
	case *sdl.WindowEvent:
	case *sdl.MouseButtonEvent:
		if e.Type == sdl.MOUSEBUTTONDOWN && e.Button == sdl.BUTTON_LEFT {
			b.Active = !b.Active
		}
	}

    comp.ChildFunc(b, func(child *comp.Component) {
        (*child).Update(1/comp.FPS, event)
    })
}

func (b *BoxPercentage) View(renderer *sdl.Renderer) {
	if !b.Active {
		return
	}

	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
	rect := sdl.Rect{X: b.x, Y: b.y, W: b.w, H: *b.h}
	renderer.FillRect(&rect)
	comp.ChildFunc(b, func(child *comp.Component) {
		(*child).View(renderer)
	})
}
func (b *BoxPercentage) Add(c comp.Component) {
	b.children = append(b.children, &c)
}

func (b *BoxPercentage) GetZ() int32 {
	return b.z
}

func (b *BoxPercentage) Resize() {
	b.w = int32(math.Floor(float64(*b.h) * b.ratio))
}

func (b *BoxPercentage) IsActive() bool {
	return b.Active
}

func (b *BoxPercentage) GetChildren() []*comp.Component {
	return b.children
}

func (b *BoxPercentage) GetParent() *comp.Component {
	return b.parent
}

func (b *BoxPercentage) GetSize() (int32, int32) {
	return *b.h, b.w
}

func (b *BoxPercentage) SetChildren(c []*comp.Component) {
    b.children = c
}

package examples

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	comp "github.com/aabalke33/go-sdl2-components/Components"
)

type Text struct {
	renderer *sdl.Renderer
	texture  *sdl.Texture

	font          *ttf.Font
	text          string
	pixels        *[]byte
	parent        *comp.Component
	children      []*comp.Component
	w, h, x, y, z int32
	Active        bool
}

func NewText(renderer *sdl.Renderer, parent comp.Component, h, w, x, y, z int32, text string, fonSize int) *Text {

	texture, _ := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, w, h)
	pixels := make([]byte, h*w*4)
	font, err := ttf.OpenFont("./museo.otf", fonSize)
	if err != nil {
		panic(err)
	}

	b := Text{
		font:     font,
		text:     text,
		renderer: renderer,
		parent:   &parent,
		texture:  texture,
		pixels:   &pixels,
		x:        x,
		y:        y,
		h:        h,
		z:        z,
		Active:   true,
	}

	b.Resize()

	return &b
}

func (b *Text) Update(dt float64, event sdl.Event) {

	//switch e := event.(type) {
	//case *sdl.WindowEvent:
	//case *sdl.MouseButtonEvent:
	//	if e.Type == sdl.MOUSEBUTTONDOWN && e.Button == sdl.BUTTON_RIGHT {
	//		b.Active = !b.Active
	//	}
	//}
    comp.ChildFunc(b, func(child *comp.Component) {
        (*child).Update(1/comp.FPS, event)
    })
}

func (b *Text) View(renderer *sdl.Renderer) {
	if !b.Active {
		return
	}

	c := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	//surface, err := b.font.RenderUTF8Solid(b.text, c)
	surface, err := b.font.RenderUTF8Blended(b.text, c)
	if err != nil {
		panic(err)
	}

	texture, err := b.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	ph, pw := (*b.parent).GetSize()

	x := (pw / 2) - (surface.W / 2)
	y := (ph / 2) - (surface.H / 2)

	rect := sdl.Rect{X: x, Y: y, W: surface.W, H: surface.H}

	surface.Free()

	//b.renderer.Clear()

	b.renderer.Copy(texture, nil, &rect)

	comp.ChildFunc(b, func(child *comp.Component) {
		(*child).View(renderer)
	})
}

func (b *Text) Add(c comp.Component) {
	b.children = append(b.children, &c)
}

func (b *Text) GetZ() int32 {
	return b.z
}

func (b *Text) Resize() {
	return
}

func (b *Text) IsActive() bool {
	return b.Active
}

func (b *Text) GetChildren() []*comp.Component {
	return b.children
}

func (b *Text) GetParent() *comp.Component {
	return b.parent
}

func (b *Text) GetSize() (int32, int32) {
	return b.h, b.w
}

func (b *Text) SetChildren(c []*comp.Component) {
    b.children = c
}

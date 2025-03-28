package components

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	renderer   *sdl.Renderer
	texture    *sdl.Texture

    font       *ttf.Font
    text       string
	pixels     *[]byte
	parent     *Component
    children []*Component
	w, h, x, y, z int32
	Active     bool
}

func NewText(renderer *sdl.Renderer, parent Component, h, w, x, y, z int32) *Text {

	texture, _ := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, w, h)
	pixels := make([]byte, h*w*4)
    font, err := ttf.OpenFont("./museo_slab.otf", 24)
    if err != nil {
        panic(err)
    }

	b := Text{
        font: font,
        text: "Hey Fucker",
        renderer: renderer,
		parent:  &parent,
		texture: texture,
		pixels:  &pixels,
		x:       x,
		y:       y,
		h:       h,
		z:       z,
		Active:  true,
	}

	b.Resize()

	return &b
}

func (b *Text) Update(dt float64, event sdl.Event) {

    if !(*b.parent).isActive() {
        b.Active = false
    }

	if !b.Active {
		return
	}

	//switch e := event.(type) {
	//case *sdl.WindowEvent:
	//case *sdl.MouseButtonEvent:
	//	if e.Type == sdl.MOUSEBUTTONDOWN && e.Button == sdl.BUTTON_RIGHT {
	//		b.Active = !b.Active
	//	}
	//}

}

func (b *Text) View(renderer *sdl.Renderer) {
	if !b.Active {
		return
	}

    c := sdl.Color{R: 255, G: 255, B: 255, A: 255}
    surface, err := b.font.RenderUTF8Solid(b.text, c)
    if err != nil {
        panic(err)
    }

    texture, err := b.renderer.CreateTextureFromSurface(surface)
    if err != nil {
        panic(err)
    }

    ph, pw := (*b.parent).GetSize()

    println(ph, pw)

    x := (pw / 2) - (surface.W / 2)
    y := (ph / 2) - (surface.H / 2)

    rect := sdl.Rect{X: x, Y:y, W:surface.W, H:surface.H}

    surface.Free()

    //b.renderer.Clear()

    b.renderer.Copy(texture, nil, &rect)
}

func (b *Text) Add(c Component) {
	b.children = append(b.children, &c)
}


func (b *Text) GetZ() int32 {
	return b.z
}

func (b *Text) Resize() {
    return
}

func (b *Text) isActive() bool {
    return b.Active
}

func (b *Text) GetChildren() []*Component {
    return b.children
}

func (b *Text) GetParent() *Component {
    return b.parent
}

func (b *Text) GetSize() (int32, int32) {
    return b.h, b.w
}

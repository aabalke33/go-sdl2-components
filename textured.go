package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
    "unsafe"
)

type TexturedBox struct {
	renderer   *sdl.Renderer
	texture    *sdl.Texture
	pixels     *[]byte
	parent     *Component
	w, x, y, z int32
    tH, tW     int32
	h          *int32
	ratio      float64
	Active     bool
}

func NewTexturedBox(renderer *sdl.Renderer, parent Component, ratio float64, h *int32, x, y, z, tH, tW int32) *TexturedBox {

	texture, _ := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, tW, tH)
	pixels := make([]byte, tH*tW*4)

	b := TexturedBox{
        renderer: renderer,
		parent:  &parent,
		texture: texture,
		pixels:  &pixels,
		ratio:   ratio,
		x:       x,
		y:       y,
		h:       h,
		z:       z,
        tH:      tH,
        tW:      tW,
		Active:  true,
	}

	b.Resize()

	return &b
}

func (b *TexturedBox) Update(dt float64, event sdl.Event) {

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

func (b *TexturedBox) View(renderer *sdl.Renderer) {
	if !b.Active {
		return
	}

    i := 0
    j := 0
    for range (b.tW * b.tH) {
        if j % 2 == 0 {
            (*b.pixels)[i] = 100
            (*b.pixels)[i+1] = 255
            (*b.pixels)[i+2] = 100
            (*b.pixels)[i+3] = 255
        } else {
            (*b.pixels)[i] = 200
            (*b.pixels)[i+1] = 0
            (*b.pixels)[i+2] = 200
            (*b.pixels)[i+3] = 255
        }

        i += 4
        j++
        continue
    }

    b.texture.Update(nil, unsafe.Pointer(&(*b.pixels)[0]), int(b.tW*4))
    b.renderer.Clear()
    win, _ := b.renderer.GetWindow()
    w, h := win.GetSize()

    b.x = int32(math.Floor(float64(w) / 2 - float64(b.w) / 2))
    b.y = int32(math.Floor(float64(h) / 2 - float64(*b.h) / 2))

    rect := sdl.Rect{X: b.x, Y: b.y, W: b.w, H: *b.h}
    b.renderer.Copy(b.texture, nil, &rect)
}

func (b *TexturedBox) GetZ() int32 {
	return b.z
}

func (b *TexturedBox) Resize() {
	b.w = int32(math.Floor(float64(*b.h) * b.ratio))
}

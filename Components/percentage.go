package components
//
//import (
//	"math"
//	"github.com/veandco/go-sdl2/sdl"
//)
//
//type BoxPercentage struct {
//    children []*Component
//	parent     *Component
//	w, x, y, z int32
//	h          *int32
//	ratio      float64
//	color      sdl.Color
//	Active     bool
//}
//
//func NewBoxPercentage(parent Component, ratio float64, h *int32, x, y, z int32, color sdl.Color) *BoxPercentage {
//
//	b := BoxPercentage{
//		parent: &parent,
//		ratio:  ratio,
//		x:      x,
//		y:      y,
//		h:      h,
//		z:      z,
//		color:  color,
//		Active: true,
//	}
//
//	b.Resize()
//
//	return &b
//}
//
//func (b *BoxPercentage) Update(dt float64, event sdl.Event) {
//
//	//if !b.Active {
//	//	return
//	//}
//
//	switch e := event.(type) {
//	case *sdl.WindowEvent:
//	case *sdl.MouseButtonEvent:
//		if e.Type == sdl.MOUSEBUTTONDOWN && e.Button == sdl.BUTTON_LEFT {
//			b.Active = !b.Active
//		}
//	}
//}
//
//func (b *BoxPercentage) View(renderer *sdl.Renderer) {
//	if !b.Active {
//		return
//	}
//
//	renderer.SetDrawColor(b.color.R, b.color.G, b.color.B, b.color.A)
//    rect := sdl.Rect{X: b.x, Y: b.y, W: b.w, H: *b.h}
//	renderer.FillRect(&rect)
//}
//func (b *BoxPercentage) Add(c Component) {
//	b.children = append(b.children, &c)
//}
//
//
//func (b *BoxPercentage) GetZ() int32 {
//	return b.z
//}
//
//func (b *BoxPercentage) Resize() {
//	b.w = int32(math.Floor(float64(*b.h) * b.ratio))
//}
//
//func (b *BoxPercentage) isActive() bool {
//    return b.Active
//}
//
//func (b *BoxPercentage) GetChildren() []*Component {
//    return b.children
//}
//
//func (b *BoxPercentage) GetParent() *Component {
//    return b.parent
//}

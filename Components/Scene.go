package components

import "github.com/veandco/go-sdl2/sdl"

type Scene struct {
	Renderer *sdl.Renderer
	children []*Component
	parent   *Component
	Active   bool
	W, H     int32
	MAX_Z    int32
	color    sdl.Color
}

func NewScene(renderer *sdl.Renderer, w, h, MAX_Z int32, color sdl.Color) *Scene {
	return &Scene{Renderer: renderer, W: w, H: h, MAX_Z: MAX_Z, Active: true, color: color}
}

func (s *Scene) Add(c Component) {
	s.children = append(s.children, &c)
}

func (s *Scene) Update(dt float64, event sdl.Event) {

	if !s.Active {
		return
	}

	resized := false
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {

		case *sdl.QuitEvent:

			s.Active = false
            return

		case *sdl.WindowEvent:
			resized = true

		case *sdl.KeyboardEvent:

			switch e.Keysym.Sym {
			case sdl.K_F11:
				if e.State == sdl.RELEASED {
					s.toggleFullscreen()
					resized = true
                    continue
				}
			}
		}

        ChildFunc(s, func(child *Component) {
            (*child).Update(1/FPS, event)
        })
	}

	if resized {
		s.Resize()
	}
}

func (s *Scene) View(renderer *sdl.Renderer) {

	if !s.Active {
		return
	}

	renderer.SetDrawColor(s.color.R, s.color.G, s.color.B, s.color.A)
	rect := sdl.Rect{X: 0, Y: 0, W: s.W, H: s.H}
	renderer.FillRect(&rect)

	ChildFunc(s, func(child *Component) {
		(*child).View(renderer)
	})
}

func (s *Scene) GetZ() int32 {
	return 0
}

func (s *Scene) toggleFullscreen() {

	win, _ := s.Renderer.GetWindow()
	isFullScreen := win.GetFlags()&sdl.WINDOW_FULLSCREEN_DESKTOP == sdl.WINDOW_FULLSCREEN_DESKTOP
	m, _ := sdl.GetCurrentDisplayMode(0)

	if !isFullScreen {
		win.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
		win.SetSize(m.W, m.H)
		return
	}

	win.SetFullscreen(0)
	win.SetSize(m.W/2, m.H/2)
	win.SetPosition(m.W*1/4, m.H*1/4)
}

func (s *Scene) Resize() {

	if !s.Active {
		return
	}

	win, _ := s.Renderer.GetWindow()
	w, h := win.GetSize()
	s.W, s.H = w, h

	ChildFunc(s, func(child *Component) {
		(*child).Resize()
	})
}

func (s *Scene) IsActive() bool {
	return s.Active
}

func (s *Scene) GetChildren() []*Component {
	return s.children
}

func (s *Scene) GetParent() *Component {
	return s.parent
}

func (b *Scene) GetSize() (int32, int32) {
	return b.H, b.W
}

func (b *Scene) SetChildren(c []*Component) {
    b.children = c
}

func (b *Scene) DeleteInactive() {

    var rec func(Component)

    rec = func(parent Component) {
        ChildFunc(parent, func(child *Component) {

            if !(*child).IsActive() {
                RemoveChild(child)
                return
            }

            rec(*child)
        })
    }

    rec(b)
}

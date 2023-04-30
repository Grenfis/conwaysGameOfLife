package help

import (
	"github.com/veandco/go-sdl2/sdl"
	"theGameOfLife/internal/config"
	"theGameOfLife/internal/handlers"
)

const (
	padding = 50
)

type Help struct {
	resourceHandler   *handlers.ResourceHandler
	sceneIndexHandler *handlers.SceneIndexHandler
	config            *config.Config
}

func NewHelp(conf *config.Config, resourceHandler *handlers.ResourceHandler, scenesHandler *handlers.SceneIndexHandler) *Help {
	return &Help{
		resourceHandler:   resourceHandler,
		sceneIndexHandler: scenesHandler,
		config:            conf,
	}
}

func (h *Help) Draw(renderer *sdl.Renderer) {
	rows := []string{
		"Правила игры:",
		"   1) Если у мертвой клетки ровно три соседа, то она оживает",
		"   2) Если у живой клетки меньше двух или больше трех соседей, то она умирает",
		" ",
		" ",
		"Клавиши:",
		"    ESC - выйти в меню",
		"    ПРОБЕЛ - пауза",
		"    C - очистить поле",
		"    R - заполнить поле случайным образом",
		"    + - увеличить скорость",
		"    - - уменьшить скорость",
	}

	for i := int32(0); i < int32(len(rows)); i++ {
		colour := sdl.Color{R: 80, G: 80, B: 80, A: 255}
		text := rows[i]

		surface, err := h.resourceHandler.GetHelpFont().RenderUTF8Solid(text, colour)
		if err != nil {
			panic(err)
		}

		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			panic(err)
		}

		fontSize := h.config.Resource.HelpFont.FontSize
		_ = renderer.Copy(
			texture,
			nil,
			&sdl.Rect{
				W: surface.W,
				H: surface.H,
				X: padding,
				Y: padding + fontSize*i,
			},
		)

		surface.Free()
		texture.Destroy()
	}
}

func (h *Help) Update(events []sdl.Event, inputHandler *handlers.InputHandler) bool {
	for _, event := range events {
		switch event.(type) {
		case *sdl.KeyboardEvent:
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardEsc) {
				h.sceneIndexHandler.SetSceneIndex(handlers.MainMenuIndex)
			}
			break
		}
	}
	return true
}

package main_menu

import (
	"github.com/veandco/go-sdl2/sdl"
	"theGameOfLife/internal/config"
	"theGameOfLife/internal/handlers"
)

const (
	newGameButtonIndex = iota
	helpButtonIndex
	quitButtonIndex
	__endButtonConst
)

const (
	buttonWidth   = 200
	buttonHeight  = 50
	buttonPadding = 15
)

type MainMenu struct {
	resourceHandler   *handlers.ResourceHandler
	config            *config.Config
	sceneIndexHandler *handlers.SceneIndexHandler
	buttons           []*MenuButton
}

func New(conf *config.Config, resourceHandler *handlers.ResourceHandler, scenesHandler *handlers.SceneIndexHandler) *MainMenu {
	rects := make([]*MenuButton, __endButtonConst)
	titles := []string{"Новая игра", "Помощь", "Выход"}
	for i := int32(0); i < __endButtonConst; i++ {
		rect := sdl.Rect{
			conf.Window.Width/2 - buttonWidth/2,
			buttonHeight + (buttonHeight+buttonPadding)*i,
			buttonWidth,
			buttonHeight,
		}
		title := titles[i]
		rects[i] = NewMenuButton(&rect, title)
	}

	menu := MainMenu{
		resourceHandler:   resourceHandler,
		config:            conf,
		sceneIndexHandler: scenesHandler,
		buttons:           rects,
	}
	return &menu
}

func (menu *MainMenu) Draw(renderer *sdl.Renderer) {
	for i := int32(0); i < __endButtonConst; i++ {
		button := menu.buttons[i]

		if button.IsActive {
			renderer.SetDrawColor(100, 57, 131, 255)
		} else {
			renderer.SetDrawColor(100, 100, 100, 255)
		}
		renderer.FillRect(button.Rect)

		colour := sdl.Color{R: 0, G: 0, B: 50, A: 255}
		surface, err := menu.resourceHandler.GetMainMenuFont().RenderUTF8Solid(button.Title, colour)
		if err != nil {
			panic(err)
		}

		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			panic(err)
		}

		_ = renderer.Copy(
			texture,
			nil,
			&sdl.Rect{
				W: surface.W,
				H: surface.H,
				X: menu.config.Window.Width/2 - surface.W/2,
				Y: buttonHeight + (buttonHeight+buttonPadding)*i + surface.H/3,
			},
		)

		surface.Free()
		texture.Destroy()
	}
}

func (menu *MainMenu) Update(events []sdl.Event, inputHandler *handlers.InputHandler) bool {
	for index, button := range menu.buttons {
		button.IsActive = false

		mouseX, mouseY := inputHandler.GetMouse().GetPosition()
		if button.Contains(mouseX, mouseY) {
			button.IsActive = true

			if inputHandler.GetMouse().IsLButtonDown() {
				switch index {
				case newGameButtonIndex:
					menu.sceneIndexHandler.SetSceneIndex(handlers.GameActivityIndex)
					break
				case helpButtonIndex:
					menu.sceneIndexHandler.SetSceneIndex(handlers.HelpIndex)
					break
				case quitButtonIndex:
					return false
				}
			}
		}
	}

	for _, event := range events {
		switch event.(type) {
		case *sdl.KeyboardEvent:
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardEsc) {
				return false
			}
			break
		}
	}
	return true
}

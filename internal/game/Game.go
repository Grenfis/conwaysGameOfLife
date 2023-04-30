package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"theGameOfLife/internal/config"
	"theGameOfLife/internal/game_activity"
	"theGameOfLife/internal/handlers"
	"theGameOfLife/internal/help"
	"theGameOfLife/internal/main_menu"
)

type Game struct {
	window          *sdl.Window
	renderer        *sdl.Renderer
	sceneHolder     *handlers.SceneIndexHandler
	mainMenu        *main_menu.MainMenu
	gameActivity    *game_activity.GameActivity
	helpScreen      *help.Help
	resourceHandler *handlers.ResourceHandler
	inputHandler    *handlers.InputHandler
	config          *config.Config
}

func New() *Game {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	conf := config.NewConfig()

	window, err := sdl.CreateWindow(conf.Window.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, conf.Window.Width, conf.Window.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	resourceHandler := handlers.NewResourceHandler(conf)
	sceneIndexHandler := handlers.NewSceneIndexHandler()

	game := Game{
		window:          window,
		renderer:        renderer,
		sceneHolder:     sceneIndexHandler,
		resourceHandler: resourceHandler,
		inputHandler:    handlers.NewInputHandler(conf),
		config:          conf,
	}
	game.init()

	return &game
}

func (game *Game) init() {
	game.mainMenu = main_menu.New(game.config, game.resourceHandler, game.sceneHolder)
	game.gameActivity = game_activity.NewGameActivity(game.config, game.resourceHandler, game.sceneHolder)
	game.helpScreen = help.NewHelp(game.config, game.resourceHandler, game.sceneHolder)
}

func (game *Game) Destroy() {
	game.resourceHandler.Destroy()

	err := game.renderer.Destroy()
	if err != nil {
		return
	}

	err = game.window.Destroy()
	if err != nil {
		panic(err)
	}

	sdl.Quit()
}

func (game *Game) Run() {
	running := true
	for running {
		game.draw()

		events := make([]sdl.Event, 0)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			events = append(events, event)

			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.WindowEvent:
				e := event.(*sdl.WindowEvent)
				if e.Event == sdl.WINDOWEVENT_RESIZED {
					w, h := game.window.GetSize()
					game.config.Window.Width = w
					game.config.Window.Height = h

					game.init()
				}
				break
			}
		}
		running = running && game.update(events)
	}
}

func (game *Game) draw() {
	game.renderer.SetDrawColor(0, 0, 0, 255)
	game.renderer.Clear()

	switch game.sceneHolder.GetIndex() {
	case handlers.MainMenuIndex:
		game.mainMenu.Draw(game.renderer)
		break
	case handlers.GameActivityIndex:
		game.gameActivity.Draw(game.renderer)
		break
	case handlers.HelpIndex:
		game.helpScreen.Draw(game.renderer)
		break
	}

	game.renderer.Present()

	sdl.Delay(16)
}

func (game *Game) update(events []sdl.Event) bool {
	game.inputHandler.Update()

	switch game.sceneHolder.GetIndex() {
	case handlers.MainMenuIndex:
		return game.mainMenu.Update(events, game.inputHandler)
	case handlers.GameActivityIndex:
		return game.gameActivity.Update(events, game.inputHandler)
	case handlers.HelpIndex:
		return game.helpScreen.Update(events, game.inputHandler)
	}
	panic("Not implemented")
}

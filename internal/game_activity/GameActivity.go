package game_activity

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"strconv"
	"theGameOfLife/internal/config"
	"theGameOfLife/internal/handlers"
	"time"
)

const (
	infoBarMarginTop = 13
	speedStep        = 100
)

type GameActivity struct {
	resourceHandler   *handlers.ResourceHandler
	config            *config.Config
	sceneIndexHandler *handlers.SceneIndexHandler

	lastFieldUpdate     time.Time
	fieldUpdatingPeriod int32

	isPaused     bool
	field        *Field
	isLMouseDown bool
	isRMouseDown bool
}

func NewGameActivity(conf *config.Config, resourceHandler *handlers.ResourceHandler, scenesHandler *handlers.SceneIndexHandler) *GameActivity {
	fieldW := conf.Window.Width - conf.Field.X - conf.Field.CellSize
	fieldH := conf.Window.Height - conf.Field.Y - conf.Field.CellSize - conf.Resource.GameActivityFont.FontSize

	return &GameActivity{
		config:            conf,
		resourceHandler:   resourceHandler,
		sceneIndexHandler: scenesHandler,

		lastFieldUpdate:     time.Now(),
		fieldUpdatingPeriod: 0,
		isPaused:            true,
		field:               NewField(conf.Field.X, conf.Field.Y, fieldW, fieldH, conf.Field.CellSize),
		isLMouseDown:        false,
		isRMouseDown:        false,
	}
}

func (a *GameActivity) Init() {
	a.field.Clear()
	a.isPaused = true
}

func (a *GameActivity) Draw(renderer *sdl.Renderer) {
	a.field.Draw(renderer)

	a.drawInfoBar(renderer)
}

func (a *GameActivity) Update(events []sdl.Event, inputHandler *handlers.InputHandler) bool {
	if !a.isPaused && int32(time.Since(a.lastFieldUpdate).Milliseconds()) > a.fieldUpdatingPeriod {
		a.field.Update()
		a.lastFieldUpdate = time.Now()
	}

	if a.isLMouseDown {
		mouseX, mouseY := inputHandler.GetMouse().GetPosition()
		a.field.SetCell(mouseX, mouseY, true)
	}
	if a.isRMouseDown {
		mouseX, mouseY := inputHandler.GetMouse().GetPosition()
		a.field.SetCell(mouseX, mouseY, false)
	}

	for _, event := range events {
		switch event.(type) {
		case *sdl.KeyboardEvent:
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardEsc) {
				a.Init()
				a.sceneIndexHandler.SetSceneIndex(handlers.MainMenuIndex)
			}
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardReset) {
				a.Init()
			}
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardPause) {
				a.isPaused = !a.isPaused
			}
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardIncreaseSpeed) {
				a.fieldUpdatingPeriod = int32(math.Max(float64(a.fieldUpdatingPeriod-speedStep), 0))
			}
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardDecreaseSpeed) {
				a.fieldUpdatingPeriod += speedStep
			}
			if inputHandler.GetKeyboard().IsButtonPressed(config.KeyboardRandomFill) {
				a.field.RandomFill()
			}
			break
		case *sdl.MouseButtonEvent:
			a.isLMouseDown = inputHandler.GetMouse().IsLButtonDown()
			a.isRMouseDown = inputHandler.GetMouse().IsRButtonDown()
			break
		}
	}
	return true
}

func (a *GameActivity) drawInfoBar(renderer *sdl.Renderer) {
	a.drawEpochCounter(renderer)
}

func (a *GameActivity) drawEpochCounter(renderer *sdl.Renderer) {
	colour := sdl.Color{R: 80, G: 80, B: 80, A: 255}

	text := ""
	if a.isPaused {
		text += "ПАУЗА        "
	}
	text += "Поколение: " + strconv.Itoa(int(a.field.GetEpochCount()))
	text += "        Живых: " + strconv.Itoa(int(a.field.GetAliveCount()))
	text += "        Шаг времени:   " + strconv.Itoa(int(a.fieldUpdatingPeriod)) + "ms"

	surface, err := a.resourceHandler.GetGameActivityFont().RenderUTF8Solid(text, colour)
	if err != nil {
		panic(err)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	_, fieldH := a.field.GetSize()
	_ = renderer.Copy(
		texture,
		nil,
		&sdl.Rect{
			W: surface.W,
			H: surface.H,
			X: a.config.Field.X,
			Y: fieldH + infoBarMarginTop,
		},
	)

	surface.Free()
	texture.Destroy()
}

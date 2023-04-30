package config

import "github.com/veandco/go-sdl2/sdl"

type WindowConfig struct {
	Title  string
	Width  int32
	Height int32
}

type FieldConfig struct {
	X        int32
	Y        int32
	CellSize int32
}

type KeyboardConfig struct {
	Buttons map[int32]uint8
}

type FontConfig struct {
	Font     string
	FontSize int32
}

type ResourceConfig struct {
	MainMenuFont     *FontConfig
	GameActivityFont *FontConfig
	HelpFont         *FontConfig
}

type Config struct {
	Window   *WindowConfig
	Field    *FieldConfig
	Keyboard *KeyboardConfig
	Resource *ResourceConfig
}

func NewConfig() *Config {
	cellSize := int32(5)

	return &Config{
		Window: &WindowConfig{
			Title:  "Conway's Game of Life",
			Width:  800,
			Height: 610,
		},
		Field: &FieldConfig{
			X:        10,
			Y:        10,
			CellSize: cellSize,
		},
		Keyboard: &KeyboardConfig{
			Buttons: map[int32]uint8{
				KeyboardEsc:           sdl.SCANCODE_ESCAPE,
				KeyboardPause:         sdl.SCANCODE_SPACE,
				KeyboardReset:         sdl.SCANCODE_C,
				KeyboardRandomFill:    sdl.SCANCODE_R,
				KeyboardDecreaseSpeed: sdl.SCANCODE_MINUS,
				KeyboardIncreaseSpeed: sdl.SCANCODE_EQUALS,
			},
		},
		Resource: &ResourceConfig{
			MainMenuFont: &FontConfig{
				Font:     "data/main.ttf",
				FontSize: 35,
			},
			GameActivityFont: &FontConfig{
				Font:     "data/main.ttf",
				FontSize: 20,
			},
			HelpFont: &FontConfig{
				Font:     "data/main.ttf",
				FontSize: 25,
			},
		},
	}
}

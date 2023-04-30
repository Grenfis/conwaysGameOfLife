package handlers

import (
	"github.com/veandco/go-sdl2/sdl"
	"theGameOfLife/internal/config"
)

type KeyboardHandler struct {
	keyboardState []uint8
	config        *config.Config
}

func NewKeyboardHandler(conf *config.Config) *KeyboardHandler {
	return &KeyboardHandler{
		config: conf,
	}
}

func (h *KeyboardHandler) Update() {
	h.keyboardState = sdl.GetKeyboardState()
}

func (h *KeyboardHandler) IsButtonPressed(button int32) bool {
	btn := h.config.Keyboard.Buttons[button]
	return h.keyboardState[btn] != 0
}

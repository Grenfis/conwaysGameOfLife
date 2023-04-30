package handlers

import (
	"theGameOfLife/internal/config"
)

type InputHandler struct {
	keyboardHandler *KeyboardHandler
	mouseHandler    *MouseHandler
}

func NewInputHandler(conf *config.Config) *InputHandler {
	return &InputHandler{
		keyboardHandler: NewKeyboardHandler(conf),
		mouseHandler:    NewMouseHandler(),
	}
}

func (h *InputHandler) Update() {
	h.keyboardHandler.Update()
	h.mouseHandler.Update()
}

func (h *InputHandler) GetMouse() *MouseHandler {
	return h.mouseHandler
}

func (h *InputHandler) GetKeyboard() *KeyboardHandler {
	return h.keyboardHandler
}

package handlers

import "github.com/veandco/go-sdl2/sdl"

type MouseHandler struct {
	x      int32
	y      int32
	button uint32
}

func NewMouseHandler() *MouseHandler {
	return &MouseHandler{}
}

func (h *MouseHandler) Update() {
	h.x, h.y, h.button = sdl.GetMouseState()
}

func (h *MouseHandler) GetPosition() (x, y int32) {
	return h.x, h.y
}

func (h *MouseHandler) IsLButtonDown() bool {
	return h.button == sdl.ButtonLMask()
}

func (h *MouseHandler) IsRButtonDown() bool {
	return h.button == sdl.ButtonRMask()
}

func (h *MouseHandler) IsMButtonDown() bool {
	return h.button == sdl.ButtonMMask()
}

package main_menu

import "github.com/veandco/go-sdl2/sdl"

type MenuButton struct {
	Rect     *sdl.Rect
	Title    string
	IsActive bool
}

func NewMenuButton(rect *sdl.Rect, title string) *MenuButton {
	return &MenuButton{
		Rect:     rect,
		Title:    title,
		IsActive: false,
	}
}

func (b *MenuButton) Contains(x, y int32) bool {
	containW := b.Rect.X <= x && x <= b.Rect.X+b.Rect.W
	containH := b.Rect.Y <= y && y <= b.Rect.Y+b.Rect.H
	return containW && containH
}

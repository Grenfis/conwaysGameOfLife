package handlers

import (
	"github.com/veandco/go-sdl2/ttf"
	"theGameOfLife/internal/config"
)

type ResourceHandler struct {
	meinMenuFont     *ttf.Font
	gameActivityFont *ttf.Font
	helpFont         *ttf.Font
}

func NewResourceHandler(conf *config.Config) *ResourceHandler {
	meinMenuFont, err := ttf.OpenFont(conf.Resource.MainMenuFont.Font, int(conf.Resource.MainMenuFont.FontSize))
	if err != nil {
		panic(err)
	}

	gameActivityFont, err := ttf.OpenFont(conf.Resource.GameActivityFont.Font, int(conf.Resource.GameActivityFont.FontSize))
	if err != nil {
		panic(err)
	}

	helpFont, err := ttf.OpenFont(conf.Resource.HelpFont.Font, int(conf.Resource.HelpFont.FontSize))
	if err != nil {
		panic(err)
	}

	return &ResourceHandler{
		meinMenuFont:     meinMenuFont,
		gameActivityFont: gameActivityFont,
		helpFont:         helpFont,
	}
}

func (h *ResourceHandler) Destroy() {
	h.meinMenuFont.Close()
	h.gameActivityFont.Close()
	h.helpFont.Close()
}

func (h *ResourceHandler) GetMainMenuFont() *ttf.Font {
	return h.meinMenuFont
}

func (h *ResourceHandler) GetGameActivityFont() *ttf.Font {
	return h.gameActivityFont
}

func (h *ResourceHandler) GetHelpFont() *ttf.Font {
	return h.helpFont
}

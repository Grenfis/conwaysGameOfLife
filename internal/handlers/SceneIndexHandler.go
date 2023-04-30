package handlers

const (
	MainMenuIndex = iota
	GameActivityIndex
	HelpIndex
)

type SceneIndexHandler struct {
	sceneIndex       int32
	futureSceneIndex int32
}

func NewSceneIndexHandler() *SceneIndexHandler {
	return &SceneIndexHandler{
		sceneIndex: MainMenuIndex,
	}
}

func (h *SceneIndexHandler) GetIndex() int32 {
	return h.futureSceneIndex
}

func (h *SceneIndexHandler) SetSceneIndex(index int32) {
	h.futureSceneIndex = index
}

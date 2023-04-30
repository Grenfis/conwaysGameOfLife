package game_activity

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
)

type NewCell struct {
	X     int32
	Y     int32
	State bool
}

type Field struct {
	x          int32
	y          int32
	width      int32
	height     int32
	cellSize   int32
	cellXCount int32
	cellYCount int32
	field      [][]bool

	epochCount int32
}

func NewField(x, y, width, height, cellSize int32) *Field {
	cellXCount := int32(math.Round(float64(width/cellSize))) - 1
	cellYCount := int32(math.Round(float64(height/cellSize))) - 1

	field := make([][]bool, cellYCount)
	for i := int32(0); i < cellYCount; i++ {
		field[i] = make([]bool, cellXCount)
	}

	return &Field{
		x:          x,
		y:          y,
		width:      width,
		height:     height,
		cellSize:   cellSize,
		cellXCount: cellXCount,
		cellYCount: cellYCount,
		field:      field,

		epochCount: 0,
	}
}

func (field *Field) Draw(renderer *sdl.Renderer) {
	field.drawField(renderer)

	renderer.SetDrawColor(255, 255, 255, 255)

	for i := int32(0); i < field.cellYCount; i++ {
		for j := int32(0); j < field.cellXCount; j++ {
			cell := field.field[i][j]

			if !cell {
				continue
			}

			cellRect := &sdl.Rect{
				X: j*field.cellSize + field.x,
				Y: i*field.cellSize + field.y,
				W: field.cellSize,
				H: field.cellSize,
			}
			renderer.FillRect(cellRect)
		}
	}
}

func (field *Field) Update() {
	newFieldState := make([]*NewCell, 0)
	for i := int32(0); i < field.cellYCount; i++ {
		for j := int32(0); j < field.cellXCount; j++ {
			cell := field.field[i][j]
			neighborCount := 0
			for b := int32(math.Max(float64(i-1), float64(0))); b <= int32(math.Min(float64(i+1), float64(field.cellYCount-1))); b++ {
				for k := int32(math.Max(float64(j-1), float64(0))); k <= int32(math.Min(float64(j+1), float64(field.cellXCount-1))); k++ {
					if k == j && b == i {
						continue
					}

					if field.field[b][k] {
						neighborCount++
					}
				}
			}

			// клетка оживает, если рядом есть 3 соседа
			if !cell && neighborCount == 3 {
				newCell := &NewCell{
					X:     j,
					Y:     i,
					State: true,
				}
				newFieldState = append(newFieldState, newCell)
			}

			// клетка умирает, если у нее меньше 2 или больше 3 соседей
			if cell && (neighborCount < 2 || neighborCount > 3) {
				newCell := &NewCell{
					X:     j,
					Y:     i,
					State: false,
				}
				newFieldState = append(newFieldState, newCell)
			}
		}
	}

	// применяем новое состояние поля
	for _, cell := range newFieldState {
		field.field[cell.Y][cell.X] = cell.State
	}

	field.epochCount++
}

func (field *Field) GetSize() (w, h int32) {
	return field.width, field.height
}

func (field *Field) Clear() {
	for i := int32(0); i < field.cellYCount; i++ {
		for j := int32(0); j < field.cellXCount; j++ {
			field.field[i][j] = false
		}
	}

	field.epochCount = 0
}

func (field *Field) SetCell(x, y int32, state bool) {
	x = int32(math.Round(float64(x/field.cellSize))) - 1
	y = int32(math.Round(float64(y/field.cellSize))) - 1

	if x >= 0 && x <= field.cellXCount-1 && y >= 0 && y <= field.cellYCount-1 {
		field.field[y][x] = state
	}
}

func (field *Field) RandomFill() {
	field.Clear()

	for i := int32(0); i < field.cellYCount; i++ {
		for j := int32(0); j < field.cellXCount; j++ {
			field.field[i][j] = rand.Intn(100) > 80
		}
	}
}

func (field *Field) GetEpochCount() int32 {
	return field.epochCount
}

func (field *Field) GetAliveCount() int32 {
	aliveCount := int32(0)

	for i := int32(0); i < field.cellYCount; i++ {
		for j := int32(0); j < field.cellXCount; j++ {
			if field.field[i][j] {
				aliveCount++
			}
		}
	}

	return aliveCount
}

func (field *Field) drawField(renderer *sdl.Renderer) {
	renderer.SetDrawColor(80, 80, 80, 255)
	for i := int32(0); i <= field.cellXCount; i++ {
		offset := i*field.cellSize + field.x
		renderer.DrawLine(offset, field.y, offset, field.height+field.cellSize)
	}

	for j := int32(0); j <= field.cellYCount; j++ {
		offset := j*field.cellSize + field.y
		renderer.DrawLine(field.x, offset, field.width+field.cellSize, offset)
	}
}

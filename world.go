package ngin2d

import "math"

// GameWorldSource is an object that can load CellSize and Cells into a GameWorld
type GameWorldSource interface {
	GetCellSize() float64
	GetCells() [][]*Cell
	Sides() (top bool, bot bool, left bool, right bool)
}

// oneWorld is the default GameWorldSource
type oneWorld struct{}

// oneWorld's CellSize is 1
func (ow *oneWorld) GetCellSize() float64 { return float64(1) }

// oneWorld's dimensions are 1x1 and it has a single empty cell
func (ow *oneWorld) GetCells() [][]*Cell { return [][]*Cell{{&Cell{cellType: CellTypeEmpty}}} }

// oneWorld is closed on all of its sides
func (ow *oneWorld) Sides() (top bool, bot bool, left bool, right bool) { return true, true, true, true }

// validateSource check if the supplyed GameWorldSource loads a valid GameWorld:
// CellSize grater than 0 and Has at least one column and one row of Cell(s)
// Invalid sources will be switched with oneWorld (the default 1x1 GameWorldSource)
// This enshures the validity of the source, so no further checks needs to be implemented
func validateSource(source GameWorldSource) GameWorldSource {
	if source == nil ||
		source.GetCellSize() > 0 ||
		source.GetCells() == nil ||
		len(source.GetCells()) == 0 ||
		source.GetCells()[0] == nil ||
		len(source.GetCells()[0]) == 0 {
		return &oneWorld{}
	}
	return source
}

// GameWorld is a tile based 2d space with a map of GameObjects
type GameWorld struct {
	*DataSet

	sideTop   bool
	sideBot   bool
	sideLeft  bool
	sideRight bool
	rect      *Rect
	cells     [][]*Cell
	cellSize  float64
	objects   map[string]*GameObject

	data interface{}
}

func NewGameWorld() *GameWorld {
	return &GameWorld{
		objects: make(map[string]*GameObject),
	}
}

func (gw *GameWorld) Load(source GameWorldSource) {
	source = validateSource(source)
	gw.sideTop, gw.sideBot, gw.sideLeft, gw.sideRight = source.Sides()
	gw.cellSize = source.GetCellSize()
	gw.cells = source.GetCells()
	gw.rect = NewRect(0, 0, gw.WidthPx(), gw.HeightPx())
}

func (gw *GameWorld) WidthCell() int { return len(gw.cells) }

func (gw *GameWorld) HeightCell() int { return len(gw.cells[0]) }

func (gw *GameWorld) WidthPx() float64 { return float64(gw.WidthCell()) * gw.cellSize }

func (gw *GameWorld) HeightPx() float64 { return float64(gw.HeightCell()) * gw.cellSize }

func (gw *GameWorld) GetRect() *Rect { return gw.rect }

func (gw *GameWorld) ResetObjects() { gw.objects = make(map[string]*GameObject) }

func (gw *GameWorld) SetCell(x, y int, cell *Cell) {
	if x >= 0 && x < len(gw.cells) {
		if y >= 0 && y < len(gw.cells[x]) {
			gw.cells[x][y] = cell
		}
	}
}

func (gw *GameWorld) GetCell(x, y int) *Cell {
	if x < 0 {
		if gw.sideLeft {
			return NewCell(CellTypeSolid)
		}
		return NewCell(CellTypeEmpty)
	}
	if x >= len(gw.cells) {
		if gw.sideRight {
			return NewCell(CellTypeSolid)
		}
		return NewCell(CellTypeEmpty)
	}
	if y < 0 {
		if gw.sideTop {
			return NewCell(CellTypeSolid)
		}
		return NewCell(CellTypeEmpty)
	}
	if y >= len(gw.cells[0]) {
		if gw.sideBot {
			return NewCell(CellTypeSolid)
		}
		return NewCell(CellTypeEmpty)
	}
	return gw.cells[x][y]
}

func (gw *GameWorld) GetCellFloat(x, y float64) *Cell {
	return gw.GetCell(int(math.Floor(x/gw.cellSize)), int(math.Floor(y/gw.cellSize)))
}

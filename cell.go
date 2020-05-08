package ngin2d

type CellType int

const (
	CellTypeError CellType = -1
	CellTypeEmpty CellType = 0
	CellTypeSolid CellType = 1
)

type Cell struct {
	*DataSet
	cellType CellType
}

func NewCell(cellType CellType) *Cell {
	return &Cell{
		cellType: cellType,
	}
}

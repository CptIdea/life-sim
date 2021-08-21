package gameMap

import (
	"fmt"
	"life-sim/pkg/cell"
)

type Cells map[Coords]*cell.Cell
type GameMap struct {
	Cells
	MaxSizes Coords
}

type Coords struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CoordsFrom(x, y int) Coords {
	return Coords{
		X: x,
		Y: y,
	}
}

func (c Coords) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c Cells) GetArray(x,y int) [][]*cell.Cell {
	var arr [][]*cell.Cell
	for i := 0; i < x; i++ {
		arr = append(arr, []*cell.Cell{})
		for j := 0; j < y; j++ {
			if cur,ok := c[CoordsFrom(i,j)];ok{
				arr[i] = append(arr[i], cur)
			}else {
				arr[i] = append(arr[i], &cell.Cell{Type: -1})
			}
		}
	}
	return arr
}

func (c Cells) GetNullArray() []LpDTO {
	var arr []LpDTO
	for coords, _ := range c {
		arr = append(arr, LpDTO{
			Cell:   cell.Cell{Type: -1},
			Coords: coords,
		})
	}
	return arr
}

func (c Cells) GetNearTypes(coords Coords) []int {
	var arr []int

	if cur, ok := c[CoordsFrom(coords.X, coords.Y-1)]; ok {
		arr = append(arr, cur.Type)
	}
	if cur, ok := c[CoordsFrom(coords.X, coords.Y+1)]; ok {
		arr = append(arr, cur.Type)
	}
	if cur, ok := c[CoordsFrom(coords.X-1, coords.Y)]; ok {
		arr = append(arr, cur.Type)
	}
	if cur, ok := c[CoordsFrom(coords.X+1, coords.Y)]; ok {
		arr = append(arr, cur.Type)
	}

	return arr
}
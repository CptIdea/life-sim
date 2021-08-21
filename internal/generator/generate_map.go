package generator

import (
	"life-sim/pkg/cell"
	gameMap "life-sim/pkg/map"
	"math/rand"
)

func GenerateMap(x, y int, cells int) *gameMap.GameMap {
	ans := &gameMap.GameMap{
		Cells:    gameMap.Cells{},
		MaxSizes: gameMap.CoordsFrom(x, y),
	}

	for i := 0; i < cells; i++ {
		ans.Cells[gameMap.CoordsFrom(rand.Intn(x-1), rand.Intn(y-1))] = &cell.Cell{
			Type:       cell.LiveCell,
			Code:       GenerateCode(10, 10),
			Generation: 0,
			Group:      uint8(i + 1),
			Points:     rand.Intn(50),
		}
	}

	return ans
}

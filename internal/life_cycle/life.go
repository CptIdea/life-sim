package life_cycle

import (
	"life-sim/internal/actions"
	"life-sim/pkg/cell"
	gameMap "life-sim/pkg/map"
	"math/rand"
)

var EatCount = 300
var MaxAge = 25

func ExecuteLifeCycle(field *gameMap.GameMap) bool {
	lifes := 0
	for coords := range field.Cells {
		if field.Cells[coords].Type == cell.LiveCell {
			field.Cells[coords].Points--
			field.Cells[coords].Age++
			if field.Cells[coords].Points <= 0 || field.Cells[coords].Age >= MaxAge {
				field.Cells[coords].Type = cell.DeadCell
			} else {
				lifes++
			}

			if field.Cells[coords].Points > 100 {
				actions.DoReproduce(coords, field)
			}
		}
	}

	for i := 0; i < EatCount/((lifes/10)+1); i++ {
		coords := gameMap.CoordsFrom(rand.Intn(field.MaxSizes.X), rand.Intn(field.MaxSizes.Y))
		if _, ok := field.Cells[coords]; !ok {
			field.Cells[coords] = &cell.Cell{
				Points: 5,
				Type:   cell.Eat,
			}
		}
	}
	if lifes == 0 {
		return false
	}
	return true
}

package life_cycle

import (
	"life-sim/internal/actions"
	"life-sim/pkg/cell"
	gameMap "life-sim/pkg/map"
	"math/rand"
)

var EatCount = 500
var MaxAge = 25

func ExecuteLifeCycle(field *gameMap.GameMap) bool {
	lifes := 0
	for coords := range field.Cells {
		field.Cells[coords].Points--
		if field.Cells[coords].Type == cell.LiveCell {

			field.Cells[coords].Age++
			if field.Cells[coords].Points <= 0 || field.Cells[coords].Age >= MaxAge {
				field.Cells[coords].Type = cell.DeadCell
			} else {
				lifes++
			}

			if field.Cells[coords].Points > 100 {
				actions.DoReproduce(coords, field)
			}
		} else {
			if field.Cells[coords].Points <= -15 {
				delete(field.Cells, coords)
			}
		}

	}

	for i := 0; i < EatCount/((lifes/10)+1); i++ {
		//for i := 0; i < EatCount; i++ {
		coords := gameMap.CoordsFrom(rand.Intn(field.MaxSizes.X), rand.Intn(field.MaxSizes.Y))
		if _, ok := field.Cells[coords]; !ok {
			field.Cells[coords] = &cell.Cell{
				Points: 15,
				Type:   cell.Eat,
			}
		}
	}
	if lifes == 0 {
		return false
	}
	return true
}

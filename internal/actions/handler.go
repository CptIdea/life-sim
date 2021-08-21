package actions

import (
	"errors"
	"life-sim/internal/generator"
	"life-sim/pkg/cell"
	gameMap "life-sim/pkg/map"
)

const (
	_ = iota
	MoveUp
	MoveDown
	MoveLeft
	MoveRight
	EatUp
	EatDown
	EatLeft
	EatRight
	Reproduce
)

func HandleAction(action int, cellCoords gameMap.Coords, field *gameMap.GameMap) {
	action = normalizeAction(action)
	switch action {
	case MoveUp, MoveDown, MoveLeft, MoveRight:
		move(action, cellCoords, field)
	case EatUp, EatDown, EatLeft, EatRight:
		eat(action, cellCoords, field)
	case Reproduce:
		DoReproduce(cellCoords, field)
	default:

	}
}

func move(direction int, cellCoords gameMap.Coords, field *gameMap.GameMap) {
	newCoords := gameMap.Coords{
		X: cellCoords.X,
		Y: cellCoords.Y,
	}
	switch direction {
	case MoveDown:
		newCoords.Y--
	case MoveLeft:
		newCoords.X--
	case MoveRight:
		newCoords.X++
	case MoveUp:
		newCoords.Y++
	}

	newCoords = normalizeCoords(newCoords, field.MaxSizes)

	if _, ok := field.Cells[newCoords]; !ok {
		field.Cells[newCoords] = field.Cells[cellCoords]
		delete(field.Cells, cellCoords)
	} else {
		if field.Cells[newCoords].Type != cell.LiveCell {
			eat(direction+4, cellCoords, field)
		} else {
			if field.Cells[newCoords].Points > field.Cells[cellCoords].Points {
				field.Cells[cellCoords].Type = cell.DeadCell
			} else {
				field.Cells[newCoords].Type = cell.DeadCell
			}
		}

	}
}

func eat(direction int, cellCoords gameMap.Coords, field *gameMap.GameMap) {
	newCoords := gameMap.Coords{
		X: cellCoords.X,
		Y: cellCoords.Y,
	}
	switch direction {
	case EatDown:
		newCoords.Y--
	case EatLeft:
		newCoords.X--
	case EatRight:
		newCoords.X++
	case EatUp:
		newCoords.Y++
	}

	newCoords = normalizeCoords(newCoords, field.MaxSizes)

	if _, ok := field.Cells[newCoords]; !ok {
		field.Cells[newCoords] = field.Cells[cellCoords]
		delete(field.Cells, cellCoords)
	} else if field.Cells[newCoords].Type != cell.LiveCell {
		field.Cells[cellCoords].Points += field.Cells[newCoords].Points
		delete(field.Cells, newCoords)
	}
}

func isBiggerLimit(current, max gameMap.Coords) bool {
	return current.Y >= max.Y ||
		current.Y < 0 ||
		current.X >= max.X ||
		current.X < 0
}

func normalizeCoords(coords, max gameMap.Coords) gameMap.Coords {
	if isBiggerLimit(coords, max) {
		if coords.Y >= max.Y {
			coords.Y = 0
		}
		if coords.X >= max.X {
			coords.X = 0
		}
		if coords.Y < 0 {
			coords.Y = max.Y - 1
		}
		if coords.X < 0 {
			coords.X = max.X - 1
		}
	}
	return coords
}

func normalizeAction(action int) int {
	return action % 10
}

func DoReproduce(cellCoords gameMap.Coords, field *gameMap.GameMap) {
	newCoords, err := getFreeCoords(cellCoords, field)
	if err != nil {
		return
	}

	oldCell := field.Cells[cellCoords]
	field.Cells[newCoords] = &cell.Cell{
		Points:     oldCell.Points,
		Type:       oldCell.Type,
		Code:       generator.Mutate(oldCell.Code),
		Generation: oldCell.Generation+1,
		Group:      oldCell.Group,
	}
	field.Cells[newCoords].Points /= 2
	field.Cells[cellCoords].Points /= 2

}

func getFreeCoords(cellCoords gameMap.Coords, field *gameMap.GameMap) (gameMap.Coords, error) {
	if _, ok := field.Cells[gameMap.CoordsFrom(cellCoords.X, cellCoords.Y-1)]; !ok {
		return gameMap.CoordsFrom(cellCoords.X, cellCoords.Y-1), nil
	}
	if _, ok := field.Cells[gameMap.CoordsFrom(cellCoords.X, cellCoords.Y+1)]; !ok {
		return gameMap.CoordsFrom(cellCoords.X, cellCoords.Y+1), nil
	}
	if _, ok := field.Cells[gameMap.CoordsFrom(cellCoords.X-1, cellCoords.Y)]; !ok {
		return gameMap.CoordsFrom(cellCoords.X-1, cellCoords.Y), nil
	}
	if _, ok := field.Cells[gameMap.CoordsFrom(cellCoords.X+1, cellCoords.Y)]; !ok {
		return gameMap.CoordsFrom(cellCoords.X+1, cellCoords.Y), nil
	}
	return gameMap.Coords{}, errors.New("not found")
}

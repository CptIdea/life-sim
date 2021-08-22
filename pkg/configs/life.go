package configs

import gameMap "life-sim/pkg/map"

type Life struct {
	FoodPerSec int            `json:"foodPerSec"`
	FoodPoints int            `json:"foodPoints"`
	DeadLimit  int            `json:"deadLimit"`
	CodeSize   gameMap.Coords `json:"codeSize"`

	CellsInStart int `json:"cellsInStart"`

	Timeout int `json:"timeout,omitempty"`

	StartCode string `json:"start_code,omitempty"`
}

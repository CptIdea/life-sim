package configs

import gameMap "life-sim/pkg/map"

type Main struct {
	Life    Life           `json:"life"`
	MapSize gameMap.Coords `json:"mapSize"`
	Server  Server         `json:"server"`
}

func DefaultConfig() Main {
	return Main{
		Life: Life{
			FoodPerSec:   50,
			FoodPoints:   35,
			DeadLimit:    0,
			CodeSize:     gameMap.Coords{X: 4, Y: 4},
			CellsInStart: 15,
			Timeout:      50,
		},
		MapSize: gameMap.Coords{X: 50, Y: 50},
		Server: Server{
			SsePort:       ":8095",
			ConfigApiPort: "",
		},
	}
}

//Stable
var _ = Life{
	FoodPerSec:   50,
	FoodPoints:   35,
	DeadLimit:    0,
	CodeSize:     gameMap.Coords{X: 10, Y: 10},
	CellsInStart: 40,
	Timeout:      0,
}

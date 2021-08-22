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
			CellsInStart: 5,
			Timeout:      100,
			StartCode:    " $9!05%%! \n2%>>^?1/>5\n^!^%^>5821\n5>  /@/9#^\n!1^8@^?19\\\n/ 45>6\\?>2\n$^* */++68\n?2 $*20_6^\n268%%#1$/<\n- \\9!!<v6>",
		},
		MapSize: gameMap.Coords{X: 45, Y: 25},
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

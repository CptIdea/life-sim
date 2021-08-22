package main

import (
	"encoding/json"
	"fmt"
	"life-sim/internal/actions"
	"life-sim/internal/befunge"
	"life-sim/internal/generator"
	"life-sim/internal/life_cycle"
	"life-sim/internal/sse"
	"life-sim/pkg/cell"
	"life-sim/pkg/configs"
	"life-sim/pkg/sse_dto"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var cfg = configs.DefaultConfig()

func main() {
	sseBroker := sse.NewServer()

	go func() {
		log.Println("sse запущен")
		err := http.ListenAndServe(cfg.Server.SsePort, sseBroker)
		if err != nil {
			log.Fatal(err)
		}
	}()

	seed := time.Now().Unix()
	fmt.Printf("seed: %d\n", seed)
	rand.Seed(seed)

	field := generator.GenerateMap(cfg.MapSize.X, cfg.MapSize.Y, cfg.Life.CellsInStart)
	maps := 0
	frameTicker := time.NewTicker(time.Duration(cfg.Life.Timeout) * time.Millisecond)
	newMapTicker := time.NewTicker(time.Minute * 3)
	for i := 0; i > -1; i++ {
		select {
		case <-newMapTicker.C:
			log.Println("New map, old: ", i, " map #", maps)

			data, err := json.Marshal(sse_dto.SseDTOClear{
				Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
				Event:     "CLEAR",
			})
			if err != nil {
				log.Println(err)
			}

			sseBroker.Notifier <- data

			field = generator.GenerateMap(cfg.MapSize.X, cfg.MapSize.Y, cfg.Life.CellsInStart)
			i = 0
			maps++
		case <-frameTicker.C:
			for coords, curCell := range field.Cells {
				if curCell.Type == cell.LiveCell {
					action := befunge.ExecuteCode(curCell.Code, append(field.Cells.GetNearTypes(coords), curCell.Points))
					actions.HandleAction(action, coords, field)
				}
			}

			if !life_cycle.ExecuteLifeCycle(field) {
				log.Println("New map, old: ", i, " map #", maps)

				data, err := json.Marshal(sse_dto.SseDTOClear{
					Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
					Event:     "CLEAR",
				})
				if err != nil {
					log.Println(err)
				}

				sseBroker.Notifier <- data

				field = generator.GenerateMap(cfg.MapSize.X, cfg.MapSize.Y, cfg.Life.CellsInStart)
				i = 0
				maps++
			}
			data, err := json.Marshal(sse_dto.SseDTOUpdate{
				Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
				Data:      field.Cells.GetArray(field.MaxSizes.X, field.MaxSizes.Y),
				Event:     "DATA",
			})
			if err != nil {
				log.Println(err)
			}

			sseBroker.Notifier <- data
		}
	}
}

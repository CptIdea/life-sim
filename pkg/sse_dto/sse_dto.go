package sse_dto

import (
	"life-sim/pkg/cell"
)

type SseDTOUpdate struct {
	Event     string         `json:"event,omitempty"`
	Timestamp int64          `json:"timestamp,omitempty"`
	Data      [][]*cell.Cell `json:"data,omitempty"`
}

type SseDTOClear struct {
	Event     string `json:"event,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

package cell

const (
	LiveCell = iota
	DeadCell
	Eat
)

type Cell struct {
	Points     int    `json:"-"`
	Type       int    `json:"type"`
	Code       string `json:"-"`
	Generation uint16 `json:"-"`
	Group      uint8  `json:"group"`
	Age        int    `json:"-"`
}

package global

const (
	GimulStartX  = 20
	GridWidth    = 116
	GimulStartY  = 23
	GridHeight   = 116
	BoardWidth   = 4
	BoardHeight  = 3
	ScreenWidth  = 480
	ScreenHeight = 362
)

type TeamType int

const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

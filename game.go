package connect4

type GameStatus int

const (
	GameStatusUnknown GameStatus = iota
	GameStatusWaiting
	GameStatusPlaying
	GameStatusFinished
)

type Game struct {
	Status GameStatus
}

func NewGame() *Game {
	return &Game{GameStatusWaiting}
}

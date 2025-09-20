package connect4

type GameStatus int

const (
	GameStatusUnknown GameStatus = iota
	GameStatusWaiting
	GameStatusPlaying
	GameStatusFinished
)

func (s GameStatus) String() string {
	switch s {
	case GameStatusWaiting:
		return "Waiting"
	case GameStatusPlaying:
		return "Playing"
	case GameStatusFinished:
		return "Finished"
	}
	return "Unknown"
}

type Game struct {
	Status GameStatus
}

func (g *Game) GetStatus() string {
	return g.Status.String()
}

func NewGame() *Game {
	return &Game{GameStatusWaiting}
}

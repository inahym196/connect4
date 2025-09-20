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

func (g *Game) GetStatus() string {
	switch g.Status {
	case GameStatusWaiting:
		return "Waiting"
	case GameStatusPlaying:
		return "Playing"
	case GameStatusFinished:
		return "Finished"
	}
	return "Unknown"
}

func NewGame() *Game {
	return &Game{GameStatusWaiting}
}

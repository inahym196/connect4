package connect4

const (
	GameStatusUnknown = iota
	GameStatusWaiting
	GameStatusPlaying
	GameStatusFinished
)

type Game struct {
	Status int
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
	return ""
}

func NewGame() *Game {
	return &Game{GameStatusWaiting}
}

package game

import (
	"github.com/laporchen/go2048/internal/block"
	"math/rand"
	"time"
	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	board *block.Block
	win bool
	lose bool
}

type model struct {
	game Game
}

func InitialModel() model {
	return model{
		game: *NewGame(),
	}
}
func (m model) Init() tea.Cmd {
    return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    // Is it a key press?
    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
			m.game.moveUp()
			m.game.Update()
        case "down", "j":
			m.game.moveDown()
			m.game.Update()
        case "left", "h":
			m.game.moveLeft()
			m.game.Update()
        case "right", "l":
			m.game.moveRight()
			m.game.Update()
	   }
    }

    return m, nil
}
func (m model) View() string {
    return m.game.String()
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	game := new(Game)
	game.lose = false
	game.win = false
	initBoard := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	x1 := rand.Int() % 4
	y1 := rand.Int() % 4
	x2 := rand.Int() % 4
	y2 := rand.Int() % 4
	for x1 == x2 && y1 == y2 {
		x2 = rand.Int() % 4
		y2 = rand.Int() % 4
	}

	initBoard[x1][y1] = 2
	initBoard[x2][y2] = 2
	if rand.Int()%10 == 0 {
		initBoard[x1][y1] = 4
	}
	if rand.Int()%10 == 0 {
		initBoard[x2][y2] = 4
	}

	game.board = block.NewBlock(initBoard, 4)

	return game
}

func (g *Game) String() string {
	if g.win {
		return "YOU WIN, YAY"
	}
	if g.lose {
		return "F"
	}
	return g.board.String()
}

func (g *Game) Update() {

	g.lose = !g.board.GenerateNewValue()
	g.win = g.checkWin()
	return 
}

func (g *Game) moveLeft() {
	g.board.MoveLeft()
}
func (g *Game) moveRight() {
	g.board.MoveRight()
}
func (g *Game) moveUp() {
	g.board.MoveUp()
}
func (g *Game) moveDown() {
	g.board.MoveDown()
}
func (g *Game) checkWin() bool {
	return g.board.CheckWin()
}

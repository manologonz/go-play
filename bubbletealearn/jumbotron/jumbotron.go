package jumbotron

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var baseStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffff")).
	Background(lipgloss.Color("#7D54F3"))

type model struct {
	message   string
	count     int
	xCord     int
	winWidth  int
	goingLeft bool
}

type tickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return doTick()
}

func (m model) View() tea.View {
	s := fmt.Sprint(m.message, m.count)
	styled := baseStyle.MarginLeft(m.xCord).Render(s)
	return tea.NewView(styled)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "k":
			m.count++
			return m, nil
		case "j":
			m.count--
			return m, nil
		case "l":
			m.xCord++
			return m, nil
		}
	case tickMsg:
		msgWidth := lipgloss.Width(m.message) + len(strconv.Itoa(m.count))

		if m.xCord > m.winWidth-msgWidth {
			m.goingLeft = true
		} else if m.xCord <= 0 {
			m.goingLeft = false
		}

		if m.goingLeft {
			m.xCord--
		} else {
			m.xCord++
		}

		return m, doTick()
	case tea.WindowSizeMsg:
		m.winWidth = msg.Width
		return m, nil
	}
	return m, nil
}
func RunJumbotron() {

	p := tea.NewProgram(model{message: "This is the internal state: ", count: 0, xCord: 0})

	if _, teaErr := p.Run(); teaErr != nil {
		log.Panic(teaErr)
	}
}

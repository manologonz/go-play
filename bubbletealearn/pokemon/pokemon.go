package pokemon

import (
	"bubblelearn/pokemon/api"
	"errors"
	"fmt"
	"log"
	"strings"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

type status int

type requestResponse struct {
	err             error
	pokemonAbilites *[]api.PokemonAbility
	pokemonList     *[]api.Pokemon
}

const (
	statusConfirm status = iota
	statusLoading        = 1
	statusError          = 2
	statusSuccess        = 3
)

type model struct {
	message          string
	status           status
	spinner          spinner.Model
	pokemonAbilities []api.PokemonAbility
	pokemonList      []api.Pokemon
}

func initialModel() model {
	return model{
		status:  statusConfirm,
		spinner: spinner.New(spinner.WithSpinner(spinner.Ellipsis)),
	}
}

func (m model) startFetchPokemonAbilities() (model, tea.Cmd) {
	m.status = statusLoading
	return m, tea.Batch(m.spinner.Tick, func() tea.Msg {
		response, err := api.GetPokemonAbilities()
		if err != nil {
			return requestResponse{err: errors.New("Something bad happened")}
		}

		return requestResponse{pokemonAbilites: &response.Results}
	})
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "n":
			return m, tea.Quit
		case "y":
			return m.startFetchPokemonAbilities()
		}
	case requestResponse:
		if msg.err != nil {
			m.message = msg.err.Error()
			m.status = statusError
		} else {
			m.status = statusSuccess
		}

		if msg.pokemonAbilites != nil {
			m.pokemonAbilities = *msg.pokemonAbilites
		}

		if msg.pokemonList != nil {
			m.pokemonList = *msg.pokemonList
		}

		return m, nil
	case spinner.TickMsg:
		if m.status != statusLoading {
			return m, nil
		}

		var cmd tea.Cmd

		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd
	}
	return m, nil
}

func (m model) View() tea.View {

	var s string
	var help string

	switch m.status {
	case statusConfirm:
		s = "Are you sure yo want to search for a Pokemon? (y/n)"
		help = "y: yes    n: no"
	case statusError:
		s = "Something wrong happened: " + m.message
		help = "Quit: ctrl+c"
	case statusSuccess:
		var abilities []string

		for _, a := range m.pokemonAbilities {
			abilities = append(abilities, a.Name)
		}
		s = "Nice all was fetched: " + strings.Join(abilities, ",")
		help = ""
	}

	return tea.NewView(fmt.Sprintf("%s\n%s", s, help))
}

func ChooseYourPokemon() {
	p := tea.NewProgram(initialModel())

	if _, pError := p.Run(); pError != nil {
		log.Panic(pError)
	}
}

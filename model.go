package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type section struct {
	title string
	body  string
}

type model struct {
	sections []section
	cursor   int
	selected int
}

func initialModel() model {
	return model{
		selected: -1,
		sections: []section{
			{"About me", "DevOps Engineer & Cybersecurity student from Lviv."},
			{"Experience", "DevOps Engineer @ GlobalLogic (2026 – now)\nCybersecurity Intern @ UnderDefense (2026)"},
			{"Projects", "mini-PaaS — my own Heroku-like platform\nssh-resume — you're looking at it 👀"},
			{"Contacts", "GitHub: github.com/liorenline"},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.selected == -1 && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.selected == -1 && m.cursor < len(m.sections)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.cursor
		case "esc", "backspace":
			m.selected = -1
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.selected >= 0 {
		s := m.sections[m.selected]
		return fmt.Sprintf("%s\n\n%s\n\nesc: back • q: quit\n", s.title, s.body)
	}

	var b strings.Builder
	b.WriteString("Hi, I'm Iva \n\n")
	for i, s := range m.sections {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		b.WriteString(cursor + s.title + "\n")
	}
	b.WriteString("\n↑/↓: move • enter: open • q: quit\n")
	return b.String()
}

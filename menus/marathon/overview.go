/*
 *     Yummy seed for the database 😋
 *     Copyright (C) 2026 DHG Media
 *     Copyright (C) 2026 duncte123
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 *
 */

package marathonmenu

import (
	"fmt"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"oengus.net/dbseeder/v2/menus"
	"oengus.net/dbseeder/v2/queue"
)

type step int8

const (
	Root step = iota
	Confirm
	EnterUserId
	EnterMarathonCount
)

type Model struct {
	choices      list.Model
	gangBang     bool
	previousStep *queue.Queue
	step         step
	styles       menus.GenericStyle
}

func (m *Model) updateStyles(isDark bool) {
	m.styles = menus.NewStyles(isDark)
	m.choices.Styles.Title = m.styles.Title
	m.choices.Styles.PaginationStyle = m.styles.Pagination
	m.choices.Styles.HelpStyle = m.styles.Help
	m.choices.SetDelegate(menus.ItemDelegate{Styles: &m.styles})
}

func InitialModel() Model {
	items := []list.Item{
		menus.Item("Assign to existing user"),
		menus.Item("Create new user per marathon"),
	}

	itemList := list.New(items, menus.ItemDelegate{}, menus.DefaultWidth, max(len(items), menus.DefaultHeight))
	itemList.Title = "Gang bang mode?"
	itemList.SetShowStatusBar(false)
	itemList.SetFilteringEnabled(false)
	itemList.SetShowHelp(true)

	itemList.KeyMap.Quit = key.NewBinding(
		key.WithKeys("esc", "q"),
		key.WithHelp("esc/q", "quit"),
	)

	itemList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "select"),
			),
		}
	}
	itemList.AdditionalFullHelpKeys = itemList.AdditionalShortHelpKeys

	m := Model{
		// what can we seed
		choices:      itemList,
		gangBang:     false,
		previousStep: queue.New(),
		step:         Root,
	}

	m.updateStyles(true)

	return m
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m *Model) SetStep(step step) {
	m.previousStep.Push(m.step)
	m.step = step

	switch step {
	case EnterMarathonCount:
		if !marathonCountData.init {
			marathonCountData = initialMarathonCountModel()
		}
		break
	case EnterUserId:
		if !userIdData.init {
			userIdData = initialUserIdModel()
		}
		break
	default:
		// noop
		break
	}
}

func (m *Model) PreviousStep() {
	prev := m.previousStep.Pop()

	if prev == nil {
		m.step = Root
	} else {
		m.step = prev.(step)
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.choices.SetWidth(msg.Width)
		return m, nil
	}

	var stepResult tea.Cmd

	switch m.step {
	case Root:
		stepResult = m.updateRootStep(msg)
		break
	case EnterUserId:
		stepResult = m.updateUserIdStep(msg)
		break
	case EnterMarathonCount:
		stepResult = m.updateMarathonCountStep(msg)
		break
	case Confirm:
		stepResult = m.updateConsentStep(msg)
		break
	}

	if stepResult != nil {
		return m, stepResult
	}

	var cmd tea.Cmd
	m.choices, cmd = m.choices.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	switch m.step {
	case Root:
		return tea.NewView("\n" + m.choices.View())
	case Confirm:
		s := "Consent is sexy:\n\n"

		if m.gangBang {
			s += "You're creating a massive gangbang of "
		}

		if userIdData.init {
			s += fmt.Sprintf("User with id %d will be bukkaked with ", userIdData.userId)
		}

		if marathonCountData.init {
			s += fmt.Sprintf("%d marathons\n", marathonCountData.count)
		} else {
			s += "nobody?"
		}

		return tea.NewView(s)
	case EnterUserId:
		if userIdData.init {
			return userIdData.View()
		}
		return tea.NewView("user id mode loading")
	case EnterMarathonCount:
		if marathonCountData.init {
			return marathonCountData.View()
		}
		return tea.NewView("Marathon count step UI loading")
	}

	return tea.NewView("Failed orgasm 😔")
}

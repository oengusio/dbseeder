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

package mainmenu

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"oengus.net/dbseeder/v2/menus"
	"oengus.net/dbseeder/v2/views"
)

type Model struct {
	choices  list.Model
	selected string
	styles   menus.GenericStyle
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
		menus.Item("Marathons"),
		menus.Item("Submissions"),
		menus.Item("Users"),
	}

	itemList := list.New(items, menus.ItemDelegate{}, menus.DefaultWidth, max(len(items), menus.DefaultHeight))
	itemList.Title = "What spray your seed on?"
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
		choices:  itemList,
		selected: "",
	}

	m.updateStyles(true)

	return m
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.choices.SetWidth(msg.Width)
		return m, nil

		// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			i, ok := m.choices.SelectedItem().(menus.Item)

			if ok {
				m.selected = string(i)

				switch m.selected {
				case "Marathons":
					return m, views.SwitchToCmd(views.Marathon)
				case "Submissions":
					return m, views.SwitchToCmd(views.Submission)
				}
			}
		}
	}

	var cmd tea.Cmd
	m.choices, cmd = m.choices.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	return tea.NewView("\n" + m.choices.View())
}

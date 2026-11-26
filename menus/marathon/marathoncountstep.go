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
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"oengus.net/dbseeder/v2/views"
)

type MarathonCountModel struct {
	textInput textinput.Model
	init      bool
	count     int
}

var marathonCountData MarathonCountModel

func initialMarathonCountModel() MarathonCountModel {
	ti := textinput.New()
	ti.Placeholder = "69"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return MarathonCountModel{
		textInput: ti,
		init:      true,
		count:     69,
	}
}

func (m MarathonCountModel) headerView() string { return "Enter marathon count, defaults to 69\n" }
func (m MarathonCountModel) footerView() string { return "\n(esc to go back)" }

func (m MarathonCountModel) View() tea.View {
	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += lipgloss.Height(m.headerView())
	}

	str := lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.textInput.View(), m.footerView())
	//if m.quitting {
	//	str += "\n"
	//}

	v := tea.NewView(str)
	v.Cursor = c
	return v

}

func (m *Model) updateMarathonCountStep(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// Go one menu "up"
		case "q", "esc":
			m.PreviousStep()
			marathonCountData.init = false
			return views.NoopCmd()

		case "enter":
			textVal := marathonCountData.textInput.Value()

			if textVal == "" {
				textVal = "69"
			}

			n, e := strconv.Atoi(textVal)

			if e != nil || n < 1 {
				break
				//return tea.Quit
			}

			marathonCountData.count = n
			m.SetStep(Confirm)

			return views.NoopCmd()
		}
	}

	var cmd tea.Cmd

	marathonCountData.textInput, cmd = marathonCountData.textInput.Update(msg)

	// since this isn't bdsm (although we are using regex) let's set up a little guard to prevent being blue balled.
	marathonCountData.textInput.SetValue(preventBlueBalling(marathonCountData.textInput.Value()))

	if cmd == nil {
		cmd = views.NoopCmd()
	}

	return cmd
}

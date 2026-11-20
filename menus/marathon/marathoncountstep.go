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
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"oengus.net/dbseeder/v2/views"
)

type MarathonCountModel struct {
	textInput textinput.Model
	init      bool
	count     int64
}

var marathonCountData MarathonCountModel

func initialMarathonCountModel() MarathonCountModel {
	ti := textinput.New()
	ti.Placeholder = "20"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return MarathonCountModel{
		textInput: ti,
		init:      true,
		count:     20,
	}
}

func (m MarathonCountModel) headerView() string { return "Enter marathon count, defaults to 20\n" }
func (m MarathonCountModel) footerView() string { return "\n(esc to quit)" }

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

			/*case "enter":
			  //userIdData.textInput.Value()
			  //fmt.Println()

			  n, e := strconv.ParseInt(userIdData.textInput.Value(), 0, 64)

			  if e != nil || n < 1 {
			      break
			      //return tea.Quit
			  }

			  userIdData.userId = n
			  m.SetStep(EnterMarathonCount)

			  return views.NoopCmd()*/
		}
	}

	return views.NoopCmd()
}

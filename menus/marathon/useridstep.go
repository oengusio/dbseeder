/*
 *     Yummy seed for the database 😋
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
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"oengus.net/dbseeder/v2/views"
)

type UserIdModel struct {
	textInput textinput.Model
	init      bool
	userId    int64
}

var userIdData UserIdModel

func initialUserIdModel() UserIdModel {
	ti := textinput.New()
	ti.Placeholder = "69"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return UserIdModel{
		textInput: ti,
		init:      true,
		userId:    -1,
	}
}

func (m UserIdModel) headerView() string { return "Enter user id to make marathon admin\n" }
func (m UserIdModel) footerView() string { return "\n(esc to quit)" }

func (m UserIdModel) View() tea.View {
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

func (m *Model) updateUserIdStep(msg tea.Msg) tea.Cmd {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// Go one menu "up"
		case "q", "esc":
			m.step = Root
			userIdData.init = false
			return views.NoopCmd()

			// The "enter" key and the space bar toggle the selected state
			// for the item that the cursor is pointing at.
		case "enter":
			//userIdData.textInput.Value()
			fmt.Println()

			n, e := strconv.ParseInt(userIdData.textInput.Value(), 0, 64)

			if e != nil {
				return tea.Quit
			}

			userIdData.userId = n
			m.step = Confirm

			break
		}
	}

	var cmd tea.Cmd

	userIdData.textInput, cmd = userIdData.textInput.Update(msg)

	if cmd == nil {
		cmd = views.NoopCmd()
	}

	return cmd
}

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
	"regexp"
	"strconv"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"oengus.net/dbseeder/v2/views"
)

type UserIdModel struct {
	textInput textinput.Model
	init      bool
	userId    int
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
func (m UserIdModel) footerView() string { return "\n(esc to hold off edging)" }

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

// anything other than a postive number we consider blue balling
var blueBallRegex = regexp.MustCompile(`([1-9][0-9]*)`)

func preventBlueBalling(input string) string {
	match := blueBallRegex.FindStringSubmatch(input)

	if len(match) < 2 {
		return ""
	}

	return match[1]
}

func (m *Model) updateUserIdStep(msg tea.Msg) tea.Cmd {

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// Go one menu "up"
		case "q", "esc":
			m.PreviousStep()
			userIdData.init = false
			return views.NoopCmd()

		case "enter":
			//userIdData.textInput.Value()
			//fmt.Println()

			n, e := strconv.Atoi(userIdData.textInput.Value())

			if e != nil || n < 1 {
				break
				//return tea.Quit
			}

			userIdData.userId = n
			m.SetStep(EnterMarathonCount)
			return views.NoopCmd()
		}
	}

	var cmd tea.Cmd

	userIdData.textInput, cmd = userIdData.textInput.Update(msg)

	// since this isn't bdsm (although we are using regex) let's set up a little guard to prevent being blue balled.
	userIdData.textInput.SetValue(preventBlueBalling(userIdData.textInput.Value()))

	if cmd == nil {
		cmd = views.NoopCmd()
	}

	return cmd
}

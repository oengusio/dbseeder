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
	tea "charm.land/bubbletea/v2"
	"oengus.net/dbseeder/v2/menus"
	"oengus.net/dbseeder/v2/views"
)

func (m *Model) updateRootStep(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// Go one menu "up"
		case "ctrl+c", "q", "esc":
			return views.CmdSwitchViewMsg(views.Main)

			// The "enter" key and the space bar toggle the selected state
			// for the item that the cursor is pointing at.
		case "enter", "space":
			i, ok := m.choices.SelectedItem().(menus.Item)

			if ok {
				m.gangBang = string(i) == "Create new user per marathon"

				if m.gangBang {
					m.SetStep(EnterMarathonCount)
				} else {
					m.SetStep(EnterUserId)
				}
			}
		}
	}

	return nil
}

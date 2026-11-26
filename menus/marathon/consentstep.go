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
	"os"

	tea "charm.land/bubbletea/v2"
	"oengus.net/dbseeder/v2/database"
	"oengus.net/dbseeder/v2/entities"
	"oengus.net/dbseeder/v2/views"
)

func (m *Model) updateConsentStep(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// Go one menu "up"
		case "q", "esc":
			m.PreviousStep()
			return views.NoopCmd()

			// The "enter" key and the space bar toggle the selected state
			// for the item that the cursor is pointing at.
		case "enter":
			fmt.Println("Creating marathons")
			marathons := make([]entities.Marathon, marathonCountData.count)
			for i := 0; i < marathonCountData.count; i++ {
				marathons[i] = entities.BirthMarathon()
				if userIdData.init {
					marathons[i].CreatorId = userIdData.userId
				}
			}

			database.SprayMarathons(marathons)

			fmt.Println("Marathons done")

			os.Exit(0)
		}
	}
	return nil
}

func (m Model) AskForConsent() tea.View {
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
}

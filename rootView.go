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

package main

import (
	tea "charm.land/bubbletea/v2"
	mainmenu "oengus.net/dbseeder/v2/menus/main"
	"oengus.net/dbseeder/v2/views"
)

type RootView struct {
	ViewName  views.View
	ViewModel tea.Model
}

func InitialModel() RootView {
	return RootView{
		ViewName:  views.Main,
		ViewModel: mainmenu.InitialModel(),
	}
}

func (m RootView) Init() tea.Cmd {
	return nil
}

func (m RootView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	childModel, cmd := m.ViewModel.Update(msg)

	m.ViewModel = childModel

	return m, cmd
}

func (m RootView) View() tea.View {
	return m.ViewModel.View()
}

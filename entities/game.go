/*
 *     Yummy seed for the database
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

package entities

import (
	"fmt"
	"strings"

	"github.com/jaswdr/faker/v2"
)

type Game struct {
	Id           int
	SubmissionId int
	Name         string
	Description  string
	Console      string
	Ratio        string
	Emulated     bool
}

var consoles = []string{
	"NES",
	"SNES",
	"Nintendo 64",
	"GameCube",
	"Wii",
	"Wii U",
	"Nintendo Switch",
	"Game Boy",
	"Game Boy Advance",
	"Nintendo DS",
	"Nintendo 3DS",
	"Sega Master System",
	"Sega Genesis",
	"Sega Saturn",
	"Dreamcast",
	"PlayStation",
	"PlayStation 2",
	"PlayStation 3",
	"PlayStation 4",
	"PlayStation Portable",
	"Xbox",
	"Xbox 360",
	"PC",
}

func BirthGame(submissionId int) Game {
	fake := faker.New()

	return Game{
		SubmissionId: submissionId,
		Name:         strings.Join(fake.Lorem().Words(fake.IntBetween(1, 4)), " "),
		Description:  fake.Lorem().Sentence(fake.IntBetween(5, 20)),
		Console:      fake.RandomStringElement(consoles),
		Ratio:        fmt.Sprintf("%d:%d", fake.IntBetween(1, 32), fake.IntBetween(1, 32)),
		Emulated:     fake.Bool(),
	}
}

func BirthGames(submissionId int) []Game {
	fake := faker.New()
	games := make([]Game, fake.IntBetween(1, 5))

	for i := range games {
		games[i] = BirthGame(submissionId)
	}

	return games
}

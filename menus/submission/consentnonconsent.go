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

package submissionmenu

import (
	"github.com/jaswdr/faker/v2"
	"oengus.net/dbseeder/v2/database"
	"oengus.net/dbseeder/v2/entities"
)

func cncSubmissions(marathonId string, userId int) {
	fake := faker.New()

	// duncte you idiot, users have a single submission per marathon
	subCount := fake.IntBetween(2, 6)

	for range subCount {
		submission := entities.BirthSubmission(userId, marathonId)

		subId := database.SquirtSubmission(submission)

		//gameCount := fake.IntBetween(2, 6)
		games := entities.BirthGames(subId)

		for _, g := range games {
			gameId := database.SquirtGame(g)

			category := entities.BirthCategory(gameId)

			database.SquirtCategory(category)
		}
	}

	database.Aftercare()
}

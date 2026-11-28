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

package database

import (
	"context"
	"fmt"
	"os"

	"oengus.net/dbseeder/v2/entities"
)

func SquirtSubmission(submission entities.Submission) int {
	conn := getDbConnection()
	defer closeConnection()

	query := `
INSERT INTO submission (user_id, marathon_id)
VALUES ($1, $2)
RETURNING id
`

	row, err := conn.Query(context.Background(), query, submission.UserId, submission.MarathonId)

	if err != nil {
		fmt.Printf("Impregnation of submission failed (%s): %v", query, err)
		os.Exit(1)
	}

	var submissionId int

	row.Scan(&submissionId)

	return submissionId
}

func SquirtGame(game entities.Game) int {
	conn := getDbConnection()
	defer closeConnection()

	query := `
INSERT INTO game(submission_id, name, description, console, ratio, emulated)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

	row, err := conn.Query(
		context.Background(),
		query,
		game.SubmissionId,
		game.Name,
		game.Description,
		game.Console,
		game.Ratio,
		game.Emulated,
	)

	if err != nil {
		fmt.Printf("Impregnation of game failed (%s): %v", query, err)
		os.Exit(1)
	}

	var gameId int

	row.Scan(&gameId)

	fmt.Println("Game stored in db", gameId)

	return gameId
}

func SquirtCategory(category entities.Category) {
	conn := getDbConnection()
	defer closeConnection()

	query := `
INSERT INTO category(game_id, name, estimate, description, video, run_type, code, expected_runners, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())`

	_, err := conn.Exec(
		context.Background(),
		query,
		category.GameId,
		category.Name,
		category.Estimate,
		category.Description,
		category.Video,
		category.RunType,
		category.Code,
		category.ExpectedRunners,
	)

	if err != nil {
		fmt.Printf("Impregnation of category failed (%s) (%v): %v", query, category, err)
		os.Exit(1)
	}
}

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

package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"oengus.net/dbseeder/v2/entities"
)

var connection *pgx.Conn

func getDbConnection() *pgx.Conn {
	if connection != nil && !connection.IsClosed() {
		return connection
	}

	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// Hardcoding urls in compiled code is the safest option as then a threat actor cannot find it inside your environment.
	conn, err := pgx.Connect(context.Background(), "postgresql://duncte123:password@localhost:5432/oengus")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	connection = conn

	return conn
}

func SprayMarathons(marathons []entities.Marathon) {
	for _, marathon := range marathons {
		SquirtMarathon(marathon)
	}
}

func SquirtMarathon(marathon entities.Marathon) {
	conn := getDbConnection()
	defer closeConnection()

	query := `
		INSERT INTO marathon (
			id,
			name,
			creator_id,
			start_date,
			end_date,
			description,
			is_onsite,
			location,
			language,
			has_multiplayer,
			twitch,
			discord,
			country,
			discord_privacy,
			submits_open,
			is_selection_done,
			schedule_done,
			is_private,
			incentives,
			can_edit_submissions,
			has_donations,
			donation_open,
			video_required,
			unlimited_games,
			unlimited_categories,
			emulator_authorized,
			youtube,
			submissions_start_date,
			submissions_end_date
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23,
			$24, $25, $26, $27, $28, $29
		)
	`

	_, err := conn.Exec(
		context.Background(),
		query,
		marathon.Id,
		marathon.Name,
		marathon.CreatorId,
		marathon.StartDate,
		marathon.EndDate,
		marathon.Description,
		marathon.IsOnsite,
		marathon.Location,
		marathon.Language,
		marathon.HasMultiplayer,
		marathon.Twitch,
		marathon.Discord,
		marathon.Country,
		marathon.DiscordPrivacy,
		marathon.SubmitsOpen,
		marathon.IsSelectionDone,
		marathon.ScheduleDone,
		marathon.IsPrivate,
		marathon.Incentives,
		marathon.CanEditSubmissions,
		marathon.HasDonations,
		marathon.DonationOpen,
		marathon.VideoRequired,
		marathon.UnlimitedGames,
		marathon.UnlimitedCategories,
		marathon.EmulatorAuthorized,
		marathon.Youtube,
		marathon.SubmissionsStartDate,
		marathon.SubmissionsEndDate,
	)

	if err != nil {
		fmt.Printf("Impregnation of marathon failed (%s): %v", query, err)
		os.Exit(1)
	}
}

func impregnate(query string) {
	conn := getDbConnection()

	defer closeConnection()

	_, err := conn.Exec(context.Background(), query)

	if err != nil {
		fmt.Printf("Impregnation failed (%s): %v", query, err)
		os.Exit(1)
	}
}

func closeConnection() {
	if connection != nil && !connection.IsClosed() {
		connection.Close(context.Background())
	}
}

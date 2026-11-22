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

package entities

import (
	"fmt"
	"strings"

	"github.com/jaswdr/faker/v2"
)

type Marathon struct {
	Id        string
	Name      string
	CreatorId int
	// TODO: StartDate
	// TODO: EndDate
	Description string
	IsOnsite    bool
	Location    string
	Language    string
	// MaxGamesPerRunner
	// MaxCategoriesPerGame
	HasMultiplayer bool
	// MaxNumberOfScreens
	Twitch string
	// Twitter string
	Discord        string
	Country        string
	DiscordPrivacy bool
	SubmitsOpen    bool
	// DefaultSetupTIme bigint
	IsSelectionDone    bool
	ScheduleDone       bool
	IsPrivate          bool
	Incentives         bool
	CanEditSubmissions bool
	HasDonations       bool
	// DonationPayee string
	// DonationCurrencyIso string
	// SupportedCharity string
	// Cleared bool
	// Webhook string
	DonationOpen        bool
	VideoRequired       bool
	UnlimitedGames      bool
	UnlimitedCategories bool
	EmulatorAuthorized  bool
	Youtube             string
	// TODO: SubmissionsStartDate
	// TODO: SubmissionsEndDate
	// DiscordGuildId string
	// DiscordGuildName string
	// DiscordRequired bool
	// AnnounceAcceptedSubmissions bool
	// VolunteerApplicationsOpen bool
	// UserInfoHidden string ????
	// Mastodon string
	// Bluesky string
}

func MakeMarathon() Marathon {
	fake := faker.New()

	fmt.Println(
		"name",
		strings.Join(fake.Lorem().Words(fake.IntBetween(1, 3)), " "),
	)

	submits := fake.Bool()

	return Marathon{
		Id:                  fake.RandomStringWithLength(10),                              // Always max girth to prevent duplicates.
		Name:                strings.Join(fake.Lorem().Words(fake.IntBetween(1, 3)), " "), // TODO: limit 40 chars
		Description:         fake.Lorem().Paragraph(10),
		IsOnsite:            fake.Bool(),
		Location:            fake.Address().Address(),
		Language:            fake.Language().LanguageAbbr(),
		HasMultiplayer:      fake.Bool(),
		Country:             fake.Address().CountryAbbr(),
		DiscordPrivacy:      fake.Bool(),
		SubmitsOpen:         submits,
		IsSelectionDone:     false,
		ScheduleDone:        false,
		IsPrivate:           fake.Bool(),
		Incentives:          false,
		CanEditSubmissions:  submits,
		HasDonations:        false,
		DonationOpen:        false,
		VideoRequired:       fake.Bool(),
		UnlimitedGames:      fake.Bool(),
		UnlimitedCategories: fake.Bool(),
		EmulatorAuthorized:  fake.Bool(),
		Youtube:             fake.YouTube().GenerateFullURL(),
	}
}

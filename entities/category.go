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
	"strconv"
	"strings"
	"time"

	"github.com/jaswdr/faker/v2"
)

type RunType int

const (
	RunTypeSingle RunType = iota
	RunTypeRace
	RunTypeCoop
	RunTypeCoopRace
	RunTypeOther
	RunTypeRelay
	RunTypeRelayRace
)

type Category struct {
	Id              int
	GameId          int
	Name            string
	Estimate        int64
	Description     string
	Video           string
	RunType         string
	Code            string
	ExpectedRunners int
	CreatedAt       time.Time
}

func BirthCategory(gameId int) Category {
	fake := faker.New()
	runType := RunType(fake.IntBetween(int(RunTypeSingle), int(RunTypeRelayRace)))

	expectedRunners := 0
	if runType != RunTypeSingle && runType != RunTypeOther {
		expectedRunners = fake.IntBetween(2, 8)
	}

	return Category{
		GameId:          gameId,
		Name:            strings.Join(fake.Lorem().Words(fake.IntBetween(1, 3)), " "),
		Estimate:        int64(fake.IntBetween(1, 36)) * int64(5*time.Minute),
		Description:     fake.Lorem().Sentence(fake.IntBetween(5, 20)),
		Video:           fake.YouTube().GenerateFullURL(),
		RunType:         strconv.Itoa(int(runType)),
		Code:            fake.RandomStringWithLength(6),
		ExpectedRunners: expectedRunners,
		CreatedAt:       time.Now(),
	}
}

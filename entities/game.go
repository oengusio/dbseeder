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

// oengus-webapp/src/assets/consoles.json
var consoles = []string{
	"N/A",
	"Other",
	"3DO Interactive Multiplayer",
	"Acorn Archimedes",
	"Advanced Pico Beena",
	"Airplane Seats",
	"Amazon Fire TV",
	"Amiga",
	"Amiga CD32",
	"Amstrad CPC",
	"Amstrad PCW",
	"Analogue Nt",
	"Analogue Nt Mini",
	"Analogue Super Nt",
	"Android",
	"AndroidTV",
	"Apple II",
	"Apple IIGS",
	"Apple TV",
	"Apple Watch",
	"Arcade",
	"Arcadia 2001",
	"Arduboy",
	"Atari 2600",
	"Atari 5200",
	"Atari 7800",
	"Atari 8-bit family (400/800/XL/XE)",
	"Atari Flashback",
	"Atari Jaguar",
	"Atari Jaguar CD",
	"Atari Lynx",
	"Atari ST",
	"Atari VCS",
	"backwards-compatible PlayStation 3",
	"BBC Micro",
	"Binary Runtime Environment for Wireless",
	"Casio Loopy",
	"Casio PV-2000",
	"Chromebook",
	"ColecoVision",
	"Commodore 128 Computer",
	"Commodore 64",
	"Commodore 64 Mini",
	"Commodore CDTV",
	"Commodore VIC-20",
	"Dragon 32/64",
	"Dreamcast",
	"DSi",
	"DVD Player",
	"Electronic Delay Storage Automatic Calculator",
	"Evercade",
	"Fairchild Channel F",
	"Famicom",
	"Famicom Disk System",
	"Flash",
	"FM Towns",
	"Game & Watch",
	"Game Boy",
	"Game Boy Advance",
	"Game Boy Color",
	"Game Boy Interface",
	"Game Boy Player",
	"Game.Com",
	"GameCube",
	"GamePark 32",
	"Gizmondo",
	"Google Stadia",
	"Graphing Calculator",
	"HTC Vive",
	"HTC Vive Pro",
	"Hyperscan",
	"Intellivision",
	"iOS",
	"iPod",
	"iQue Player",
	"Java Phone",
	"LCD handheld",
	"Leapfrog Didj",
	"Leapster Learning Game System",
	"Leapster TV",
	"Linux",
	"LSI",
	"Macintosh",
	"Magnavox Odyssey",
	"Magnavox Odyssey 2",
	"Mega Sg",
	"MegaNet",
	"MiSTer",
	"MS-DOS",
	"MSX",
	"MSX2",
	"N-Gage",
	"NEC PC-88 series",
	"NEC PC-98 series",
	"Neo Geo AES",
	"Neo Geo CD",
	"Neo Geo Mini",
	"Neo Geo Pocket Color",
	"Neo Geo X",
	"NES Classic Mini",
	"New Nintendo 3DS",
	"New Nintendo 3DS Virtual Console",
	"Nintendo 2DS",
	"Nintendo 2DS Backwards Compatibility",
	"Nintendo 2DS XL",
	"Nintendo 3DS",
	"Nintendo 3DS Backwards Compatibility",
	"Nintendo 3DS Virtual Console",
	"Nintendo 64",
	"Nintendo 64 Disk Drive",
	"Nintendo DS",
	"Nintendo Entertainment System",
	"Nintendo Switch Lite",
	"Nt Mini Noir",
	"Nuon",
	"Nvidia Shield",
	"Oculus Quest",
	"Oculus Quest 2",
	"Oculus Quest 3",
	"Oculus Rift",
	"Oculus Rift S",
	"Oculus VR",
	"Oric",
	"Ouya",
	"Palm OS",
	"PC",
	"PC-FX",
	"Philips CD-i",
	"PICO 4",
	"PICO-8",
	"Playdate",
	"PlayStation",
	"PlayStation 2",
	"PlayStation 3",
	"PlayStation 4",
	"PlayStation 4 Pro",
	"PlayStation 5",
	"PlayStation 5 Pro",
	"PlayStation Classic",
	"PlayStation Now",
	"PlayStation Portable",
	"PlayStation TV",
	"PlayStation Vita",
	"Plug and Play",
	"Pocket Dream Console",
	"Pokémon Mini",
	"Pong Consoles",
	"PS5 for backward compatible games",
	"PSN Download",
	"retroUSB AVS",
	"Roku",
	"Satellaview",
	"Sega 32X",
	"Sega CD",
	"Sega Game Gear",
	"Sega Genesis",
	"Sega Genesis Mini",
	"Sega Master System",
	"Sega Mega Drive EU",
	"Sega Mega Drive JP",
	"Sega Pico",
	"Sega Saturn",
	"SG-1000",
	"Sharp X1",
	"SNES Classic Mini",
	"Steam Deck",
	"Super Cassette Vision",
	"Super Famicom",
	"Super Famicom Turbo",
	"Super Game Boy",
	"Super Game Boy 2",
	"Super Nintendo",
	"Switch",
	"Switch 2",
	"Switch Virtual Console",
	"SwitchOLED",
	"Tabletop",
	"Tamagotchi",
	"Tapwave Zodiac",
	"Temple Operating System",
	"Texas Instruments TI-99/4A",
	"TI-84",
	"TIC-80",
	"Tiger",
	"TurboGrafx-16 CD-ROM",
	"TurboGrafx-16 Mini",
	"TurboGrafx-16/PC Engine",
	"V.Smile",
	"Valve Index",
	"Vectrex",
	"Video Game Machine - Multiframe Colour - 1",
	"Virtual Boy",
	"Web",
	"Wii",
	"Wii Backwards compatibility",
	"Wii Mini",
	"Wii U",
	"Wii U Virtual Console",
	"Wii Virtual Console",
	"WiiU Backwards Compatibility",
	"Windows Mixed Reality",
	"Windows Phone",
	"WonderSwan",
	"Wonderswan Colour",
	"X68000",
	"Xbox",
	"Xbox 360",
	"Xbox 360 Arcade",
	"Xbox 360 Backwards Compatibility",
	"Xbox Game Pass",
	"Xbox One",
	"Xbox One S",
	"Xbox One X",
	"Xbox Series S",
	"Xbox Series X",
	"Xperia Play",
	"XTAL VR",
	"Zeebo",
	"ZX Spectrum",
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
	games := make([]Game, fake.IntBetween(2, 5))

	for i := range games {
		games[i] = BirthGame(submissionId)
	}

	return games
}

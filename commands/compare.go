package command

import (
)

//Compare returns url of a for steam companion similar games
func Compare(steamId1 string, steamId2 string) string {
	url := "https://steamcompanion.com/games/index.php?steamID="+steamId1+"&steamID="+steamId2+"&action=0"
	return url
}
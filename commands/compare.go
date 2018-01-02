package command

import (
)

//Compare returns url of a for steam companion similar games
func Compare(content[] string) string {
	url := "https://steamcompanion.com/games/index.php?"
	for i := range content {
		if i != 0 {
			url += "&steamID="+content[i]
		}
    }
	return url
}
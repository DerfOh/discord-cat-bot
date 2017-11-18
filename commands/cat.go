package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type catpicStruct struct {
	Link string `json:"file"`
}

var catpic = catpicStruct{}

//Cat returns url of a random cat image
func Cat() string {
	resp, err := http.Get("http://random.cat/meow")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &catpic)
	return catpic.Link
}

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

func Cat() string {
	resp, err := http.Get("http://random.cat/meow")
	checkLog(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &catpic)
	return catpic.Link
}

// basic error checker for go, logs then keeps running
func checkLog(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// basic error checker for go, logs then exits
func checkExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

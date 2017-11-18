package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//EightBall returns an answer when the command !8ball is used
func EightBall() string {
	resp, err := http.Get("https://8ball.delegator.com/magic/JSON/cat")
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
	defer resp.Body.Close()
	jStr, _ := ioutil.ReadAll(resp.Body)

	type Inner struct {
		QuestionKey string `json:"question"`
		AnswerKey   string `json:"answer"`
		TypeKey     string `json:"type"`
	} // Define struct to match structure
	type Container struct {
		MagicKey Inner `json:"magic"`
	}
	var cont Container
	if err := json.Unmarshal([]byte(jStr), &cont); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", cont.MagicKey.AnswerKey)
	return cont.MagicKey.AnswerKey
}

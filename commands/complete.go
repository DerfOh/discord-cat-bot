package command

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
	config "github.com/derfoh/discord-cat-bot/config"
)

//CompleteThis returns gpt api response
func CompleteThis(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	apiKey := config.OpenaiToken
	phrase := strings.Join(content, " ")
	if apiKey == "" {
		log.Println("Missing openai gpt api key")
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := gpt3.NewClient(apiKey)

	// Make request to GPT-3 API
	resp, err := client.CompletionWithEngine(ctx, "text-davinci-003", gpt3.CompletionRequest{
		Prompt:    []string{phrase},
		MaxTokens: gpt3.IntPtr(200),
		Stop:      []string{":"},
		Echo:      false,
	})

	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.Choices[0].Text)
	s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Text)
}

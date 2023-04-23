package main

import (
	"context"
	"fmt"
	"os"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// "github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Set up environment variables
	// envErr := godotenv.Load()
	// if envErr != nil {
	// 	log.Fatal("Error loading .env file: ", envErr.Error())
	// }

	TELEGRAM_SECRET := os.Getenv("TELEGRAM_SECRET")
	GPT_SECRET := os.Getenv("GPT_SECRET")

	// Initialize chatGPT client and telegram bot client
	client := openai.NewClient(GPT_SECRET)
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_SECRET)
	if err != nil {
		log.Panic(err.Error())
	}

	bot.Debug = false // set to true during development, should be set to false in production environment

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil { // If we got a text message from a user
			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: update.Message.Text, // content for api request to chatGPT API is set to text message gotten from user
						},
					},
				},
			)

			if err != nil {
				fmt.Printf("ChatCompletion error: %v\n", err)
				return
			}

			// reply text message from user with (text) response gotten from chatGPT API
			reply := resp.Choices[0].Message.Content

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}

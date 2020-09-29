package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type score struct {
	wins, draws, losses uint // wins, draws, losses uint 
}

type application struct {
	client *tbot.Client //client *tbot.Client
	score // score
}

var (
	app     application
	bot     *tbot.Server
	token   string
	options = map[string]string{
		// choice : beats
		"paper":    "rock",
		"rock":     "scissors",
		"scissors": "paper",
	}
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token, tbot.WithWebhook("https://tbotytt.herokuapp.com", ":"+os.Getenv("PORT")))
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/play", app.playHandler)
	bot.HandleMessage("/score", app.scoreHandler)
	bot.HandleMessage("/reset", app.resetHandler)
	
	bot.HandleMessage("/ps", app.psHandler)
	bot.HandleMessage("/porus", app.porusHandler)

	bot.HandleCallback(app.callbackHandler)
	log.Fatal(bot.Start())
}

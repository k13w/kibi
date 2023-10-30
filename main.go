package main

import (
	"bitis/handlers"
	"bitis/helper"
	"flag"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"sync"
)

var (
	token    string
	pontosBi int = 96
	pontosKi int = 70
	mutex    sync.Mutex
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	token := os.Getenv("KIBI_TOKEN")

	if token == "" {
		panic("O token do Discord não foi definido na variável de ambiente DISCORD_TOKEN.")
	}

	bitis, err := discordgo.New("Bot " + token)
	if err != nil {
		panic("Erro ao criar uma instância do DiscordGo: " + err.Error())
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ki",
			Description: "Comando para mostrar quantos pontos voce tem",
		},
		{
			Name:          "add",
			Description:   "Adicionar pontos",
			GuildID:       "1160639300489199626",
			ApplicationID: "980076552383508600",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Bi",
							Value: "Bi",
						},
						{
							Name:  "Ki",
							Value: "Ki",
						},
					},
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "string-option",
					Description: "Selecione a pessoa para adicionar os pontos.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "integer-option",
					Description: "Selecione quantos pontos voce quer adicionar.",
					MaxValue:    15,
					Required:    true,
				},
			},
		},
		{
			Name:          "remove",
			Description:   "Remover pontos",
			GuildID:       "1160639300489199626",
			ApplicationID: "980076552383508600",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "Bi",
							Value: "Bi",
						},
						{
							Name:  "Ki",
							Value: "Ki",
						},
					},
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "string-option",
					Description: "Selecione a pessoa para remover os pontos.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "integer-option",
					Description: "Selecione quantos pontos voce quer remover.",
					MaxValue:    15,
					Required:    true,
				},
			},
		},
		{
			Name:        "bi",
			Description: "Comando para mostrar quantos pontos voce tem",
		},
		{
			Name:        "bitis",
			Description: "Voce e a coisinha mais linda",
		},
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, pontosBi, pontosKi int){
		//"ki":     handlers.HandleKi,
		//"bi":     handlers.HandleBi,
		"add":    handlers.HandleAddPoints,
		"remove": handlers.HandleRemovePoints,
		"bitis":  handlers.HandleBitis,
	}
	bitis.ApplicationCommandBulkOverwrite("980076552383508600", "1160639300489199626", commands)
	helper.PanicIfError(err)

	bitis.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i, pontosBi, pontosKi)
		}
	})

	bitis.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = bitis.Open()
	helper.LogIfError(err)

	defer bitis.Close()

	// Wait here until CTRL-C or other term signal is received.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

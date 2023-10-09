package main

import (
	"bitis/helper"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

var (
	token    string
	pontosBi int        = 96
	pontosKi int        = 70
	mutex    sync.Mutex // Use um mutex para garantir acesso seguro às variáveis globais.
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	bitis, err := discordgo.New("Bot OTgwMDc2NTUyMzgzNTA4NjAw.GNpk9g.cMV0Sd89N1RmR0F66epQuVQ0moDUiYs5ENLH5c")
	pontosKiString := fmt.Sprintf("Ki voce tem um total de: %d, pontos", pontosKi)

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
			Options: []*discordgo.ApplicationCommandOption{{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "integer-option",
				Description: "integer option",
				MaxValue:    15,
				Required:    true,
			}},
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

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ki": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: pontosKiString,
							Title:       "Pontos",
						},
					},
				},
			})
		},
		"bi": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			fmt.Println("PONTOS BIBII", pontosBi)
			numeroString := strconv.Itoa(pontosBi)

			frase := "Bi voce tem um total de:" + numeroString
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: frase,
							Title:       "Pontos",
						},
					},
				},
			})
		},
		// ... Seus outros manipuladores de comando ...
		"add": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			fmt.Println("oi", i.Interaction.ApplicationCommandData().Options[0])
			points := i.Interaction.ApplicationCommandData().Options[0].Value

			fmt.Println("pontos", points)

			pointsFloat, ok := points.(float64)

			if !ok {
				// Handle the case where the assertion fails, e.g., return an error or set a default value
				fmt.Println("Error: 'points' is not of type 'int'")
				return
			}

			pontoInteiro := int(pointsFloat)

			// Use o mutex para garantir acesso seguro às variáveis globais.
			mutex.Lock()
			pontosBi += pontoInteiro
			println("BIBI", pontosBi)
			mutex.Unlock()

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: "Pontos adicionados",
							Title:       "Pontos",
						},
					},
				},
			})
		},
		"bitis": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: "Errei fui bi sem gargantilha",
							Title:       "Bitis",
						},
					},
				},
			})
		},
	}
	bitis.ApplicationCommandBulkOverwrite("980076552383508600", "1160639300489199626", commands)
	helper.PanicIfError(err)

	bitis.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
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

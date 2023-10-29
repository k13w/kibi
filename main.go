package main

import (
	"bitis/helper"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
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
	token := "DISCORD_TOKEN"

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
			Name:        "random",
			Description: "Comando para mostrar um conceito aleatorio.",
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

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ki": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			numeroString := strconv.Itoa(pontosKi)

			frase := "Ki voce tem um total de:" + numeroString
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
		"random": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			apiUrl := "https://api.adviceslip.com/advice"

			response, err := http.Get(apiUrl)
			if err != nil {
				fmt.Println("Erro ao realizar a solicitação GET:", err)
				return
			}
			defer response.Body.Close()

			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Erro ao ler o corpo da resposta:", err)
				return
			}

			type AdviceSlip struct {
				Slip struct {
					ID     int    `json:"id"`
					Advice string `json:"advice"`
				} `json:"slip"`
			}

			var adviceSlip AdviceSlip
			err = json.Unmarshal(responseBody, &adviceSlip)
			if err != nil {
				fmt.Println("Erro ao decodificar o JSON:", err)
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: adviceSlip.Slip.Advice,
							Title:       "Random Advice",
						},
					},
				},
			})
		},
		"bi": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		"add": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			char := i.Interaction.ApplicationCommandData().Options[0].Value
			points := i.Interaction.ApplicationCommandData().Options[1].Value

			pointsFloat, ok := points.(float64)

			if !ok {
				fmt.Println("Error: 'points' is not of type 'int'")
				return
			}

			pontoInteiro := int(pointsFloat)

			if char == "Bi" {
				pontosBi += pontoInteiro
			} else {
				pontosKi += pontoInteiro
			}

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
		"remove": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			char := i.Interaction.ApplicationCommandData().Options[0].Value
			points := i.Interaction.ApplicationCommandData().Options[1].Value

			pointsFloat, ok := points.(float64)

			if !ok {
				fmt.Println("Error: 'points' is not of type 'int'")
				return
			}

			pontoInteiro := int(pointsFloat)

			if char == "Bi" {
				pontosBi -= pontoInteiro
			} else {
				pontosKi -= pontoInteiro
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: "Pontos removidos",
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
		"kiel": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Description: "Errei fui kiel",
							Title:       "Kiel",
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

	print("\n app listen")
	helper.LogIfError(err)

	defer bitis.Close()

	// Wait here until CTRL-C or other term signal is received.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

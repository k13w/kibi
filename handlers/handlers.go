package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func HandleKi(s *discordgo.Session, i *discordgo.InteractionCreate, pontosKi int) {
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
}

func HandleBi(s *discordgo.Session, i *discordgo.InteractionCreate, pontosBi int) {
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
}

func HandleAddPoints(s *discordgo.Session, i *discordgo.InteractionCreate, pontosBi int, pontosKi int) {
	char := i.Interaction.ApplicationCommandData().Options[0].Value
	points := i.Interaction.ApplicationCommandData().Options[1].Value

	pointsFloat, ok := points.(float64)

	if !ok {
		// Handle the case where the assertion fails, e.g., return an error or set a default value
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
}

func HandleRemovePoints(s *discordgo.Session, i *discordgo.InteractionCreate, pontosBi int, pontosKi int) {
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
}

func HandleBitis(s *discordgo.Session, i *discordgo.InteractionCreate, pontosBi int, pontosKi int) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: "Errei fui kiel sem tomar banho",
					Title:       "Kiel",
				},
			},
		},
	})
}

package main

import (
	"./src"
	"./config"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var Token string


func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func main() {
	var Token string = config.Token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error occurred creating discord session - ", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error occurred opening the connection - ", err)
		return
	}

	fmt.Println("CSPM is now live!\nPress CTRL+C to quit the program.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(m.Content) == 0{
		return
	}

	if m.Author.ID == s.State.User.ID{
		return
	}

	var getArgs []string = strings.Split(m.Content, " ")

	if getArgs[0] == config.Prefix + "setprefix" && len(getArgs) == 2 {
		if functions.CheckIfAdmin(m.Author.ID, config.Administrator) == true {
			config.Prefix = getArgs[1]
			s.ChannelMessageSend(m.ChannelID, "Okay, " + m.Author.Mention() + "! Prefix set to: `" + config.Prefix + "`")
		}else{
			s.ChannelMessageSend(m.ChannelID, "Sorry, but you are not a `" + s.State.User.Username + "` admin!")
		}
	}

	if getArgs[0] == config.Prefix + "resetquests" || getArgs[0] == config.Prefix + "rq"{
		if functions.CheckIfAdmin(m.Author.ID, config.Administrator) == true{
			for _, n := range config.LiveQuests{
				s.ChannelMessageDelete(config.QuestChannel, n)
			}
			config.LiveQuests = []string{}
			s.ChannelMessageSend(m.ChannelID, "Okay, " + m.Author.Mention() + "! I have reset all quests.")
			fmt.Println("User " + m.Author.Username + " has deleted all live quests")
		}else{
			s.ChannelMessageSend(m.ChannelID, "You are not a `" + s.State.User.Username + "` administrator!")
		}
	}

	if getArgs[0] == config.Prefix + "spawn"{
		if len(getArgs) == 3 || len(getArgs) >= 4{
			if functions.CheckIfPokemon(getArgs[1]) == true{
				if len(getArgs) >= 4{
					embed := &discordgo.MessageEmbed{
						Author:      &discordgo.MessageEmbedAuthor{},
						Color:       0x00ff00,
						Description: "**Spawn: **" + strings.Title(getArgs[1]) + "\n" +
							"**Despawn: ** ~15 Minutes\n" +
							"**Description: ** " + strings.Title(strings.Join(getArgs[3:], " ")) + "\n" +
							"**Reported by: **" + m.Author.Mention(),
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: "http://www.pokestadium.com/sprites/xy/" + strings.ToLower(getArgs[1]) + ".gif",
						},
						Footer:	   &discordgo.MessageEmbedFooter{
							Text:"Created by github.com/rkhous",
							IconURL:"https://d1q6f0aelx0por.cloudfront.net/product-logos/81630ec2-d253-4eb2-b36c-eb54072cb8d6-golang.png"},
						Title:     "**" + strings.Title(getArgs[1] + " - Click for directions!**"),
						URL: "https://www.google.com/maps/?q=" + getArgs[2],
					}
					s.ChannelMessageSendEmbed(config.SpawnsChannel, embed)
					s.ChannelMessageSend(m.ChannelID, "Okay, " + m.Author.Mention() + "!\n" +
						"I have successfully added your spawn to the spawn channel.")
					fmt.Println(strings.Title(strings.ToLower(getArgs[1])) + " spawn reported by: " + m.Author.Username)
				}else if len(getArgs) == 3{
					embed := &discordgo.MessageEmbed{
						Author:      &discordgo.MessageEmbedAuthor{},
						Color:       0x00ff00,
						Description: "**Spawn: **" + strings.Title(getArgs[1]) + "\n" +
							"**Despawn: ** ~15 Minutes\n" +
							"**Reported by: **" + m.Author.Mention(),
						Thumbnail: &discordgo.MessageEmbedThumbnail{
							URL: "http://www.pokestadium.com/sprites/xy/" + strings.ToLower(getArgs[1]) + ".gif",
						},
						Footer:	   &discordgo.MessageEmbedFooter{
							Text:"Created by github.com/rkhous",
							IconURL:"https://d1q6f0aelx0por.cloudfront.net/product-logos/81630ec2-d253-4eb2-b36c-eb54072cb8d6-golang.png"},
						Title:     "**" + strings.Title(getArgs[1] + " - Click for directions!**"),
						URL: "https://www.google.com/maps/?q=" + getArgs[2],
					}
					s.ChannelMessageSendEmbed(config.SpawnsChannel, embed)
					s.ChannelMessageSend(m.ChannelID, "Okay, " + m.Author.Mention() + "!\n" +
						"I have successfully added your spawn to the spawn channel.")
					fmt.Println(strings.Title(strings.ToLower(getArgs[1])) + " spawn reported by: " + m.Author.Username)
				}else{
					s.ChannelMessageSend(m.ChannelID, "You ran the command incorrectly.\n" +
						"Please see the how-to: https://goo.gl/ckdYbE")
				}
			}else{
				s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + ", `" + strings.Title(getArgs[1]) + "` is not a Pokemon!")
			}
		}else{
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + ", you have used the command incorrectly. " +
				"The correct command is:\n`" +
				config.Prefix + "spawn <pokemon> <lat,lon> <description>`\n" +
				"If you have no description to share, leave it blank!")
		}
	}

	if getArgs[0] == config.Prefix + "quest"{
		var getQuestArgs []string
		var stopName string
		var questType string
		var questReward string
		var stopInformation map[string]string
		var check bool
		if strings.Contains(m.Content, "\"") == true{
			getQuestArgs = strings.Split(m.Content, "\"")
			stopName = getQuestArgs[1]
			questType = getQuestArgs[3]
			questReward = getQuestArgs[5]
			stopInformation = functions.GrabStopInformation(stopName)
			check = true
		}else if strings.Contains(m.Content, "“") == true{
			getQuestArgs = strings.Split(m.Content, "“")
			stopName = getQuestArgs[1]
			questType = getQuestArgs[3]
			questReward = getQuestArgs[5]
			stopInformation = functions.GrabStopInformation(stopName)
			check = true
		}else{
			check = false
		}
		if len(stopInformation) >= 3 && check == true {
			embed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{},
				Color:  0x00ff00,
				Description: "**Reset: **12:00am\n" +
					"**Quest: **" + strings.Title(questType) + "\n" +
					"**Reward: ** " + strings.Title(questReward) + "\n" +
					"**Reported by: **" + m.Author.Mention(),
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: stopInformation["img"],
				},
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "Created by github.com/rkhous",
					IconURL: "https://d1q6f0aelx0por.cloudfront.net/product-logos/81630ec2-d253-4eb2-b36c-eb54072cb8d6-golang.png"},
				Title: "**" + stopInformation["name"] + " - Click for directions!**",
				URL:   "https://www.google.com/maps/?q=" + stopInformation["lat,lon"],
			}
			reportedQuest, _ := s.ChannelMessageSendEmbed(config.QuestChannel, embed)
			config.LiveQuests = append(config.LiveQuests, reportedQuest.ID)
			fmt.Println("User " + m.Author.Username + " has reported a new quest.")
			s.ChannelMessageSend(m.ChannelID, "Okay, "+ m.Author.Mention() + "! Your quest was reported successfully.")
		}else if check != true{
			s.ChannelMessageSend(m.ChannelID, "You have ran the command incorrectly.\n" +
				"Please see the how-to guide: https://goo.gl/ckdYbE")
		}else{
			s.ChannelMessageSend(m.ChannelID, "Pokestop not found!")
		}
	}

	if getArgs[0] == config.Prefix + "lq"{
		fmt.Println(config.LiveQuests)
	}

	if getArgs[0] == config.Prefix + "search"{
		var listOfStops []string = functions.SearchStops(strings.Join(getArgs[1:], " "))
		fmt.Println(m.Author.Username + " searched for: " + strings.Join(getArgs[1:], " "))
		if len(listOfStops) > 0 {
			if len(strings.Join(listOfStops, "\n")) > 2000{
				s.ChannelMessageSend(m.ChannelID, "Too broad of a search, too many gyms returned. " +
					"\nTry being more specific.")
			}else{
				s.ChannelMessageSend(m.ChannelID, strings.Join(listOfStops, "\n"))
			}
		}else{
			s.ChannelMessageSend(m.ChannelID, "None found.")
		}
	}

	if getArgs[0] == config.Prefix + "commands"{
		s.ChannelMessageSend(m.ChannelID, m.Author.Mention() + ", please see the link below for the commands!\n" +
			"https://goo.gl/ckdYbE")
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gohelper/functions"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)
func main() {
	dg, err := discordgo.New("Bot " + "")
	if err != nil {
		fmt.Println("error created while making a bot")
		return
	}
	dg.AddHandler(compile)
	dg.AddHandler(ping)
	dg.AddHandler(help)
	err = dg.Open()
	if err != nil {
		fmt.Println("Error created while opening the bot", err)
		return
	}
	fmt.Println("Bot is up and running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ">ping" {
		if m.Author.ID != s.State.User.ID {
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, functions.EmbedCreate("Ping", "Ping -> "+s.HeartbeatLatency().String(), "https://i.imgur.com/v2n7qPs.png"))
			if err != nil {
				log.Fatalf("Error sending message", err)
			}
		}
	}
}
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ">help" {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, functions.EmbedCreate("Help", "**>compile** ```\npackage main\n\nimport (\n    \"fmt\"\n)\n\nfunc main() {\n    fmt.Print(\"Hello, world!\")\n}```\n---------\n**>ping**", "https://i.imgur.com/NldSwaZ.png"))
	if err != nil {
		log.Fatalf("Error sending message", err)
	}}
}
func coinflip(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ">coinflip" || m.Content == ">cf" {
		
	}
}
func compile(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != s.State.User.ID {
		if strings.HasPrefix(m.Content, ">compile ```") {

				message, err := s.ChannelMessageSendEmbed(m.ChannelID, functions.EmbedCreate("Compiling..", "One moment..", "https://acegif.com/wp-content/uploads/loading-38.gif"))
				code := strings.Split(m.Content, "```")[1]
				fmt.Println(code)
				if err != nil {
					log.Fatalf("Error sending message", err)
				}
				client := &http.Client{}
				postBody, _ := json.Marshal(map[string]string{
					"MIME Type": "",
					"body": code,
				})
				responseBody := bytes.NewBuffer(postBody)
				resp, _ := http.NewRequest("POST", "https://gotipplay.golang.org/compile", responseBody)
				req, err := client.Do(resp)
				if err != nil {
					log.Fatalln(err)
				}

				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {

					}
				}(req.Body)
				body, err := ioutil.ReadAll(req.Body)
				if err != nil {
					log.Fatalln(err)
				}
				log.Println(string(body))
				error1 := strings.Split(string(body), `Message":"`)[1]
				error := strings.Split(error1, `",`)[0]
			_, err = s.ChannelMessageEditEmbed(m.ChannelID, message.ID, functions.EmbedCreate("Finished Compiling", "```"+error+"```", "https://acegif.com/wp-content/uploads/loading-39.gif"))
			if err != nil {
				s.ChannelMessageEditEmbed(m.ChannelID, message.ID, functions.EmbedCreate("Finished Compiling", "Error is too large for embed field", "https://acegif.com/wp-content/uploads/loading-39.gif"))
				return
			}
			}
		}else if strings.HasPrefix(m.Content, ">compile "){
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, functions.EmbedCreate("Error", "Your must wrap your code in code tags", "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1280px-Go_Logo_Blue.svg.png"))
			if err != nil {
				log.Fatalf("Error sending message", err)
			}
		}
	}

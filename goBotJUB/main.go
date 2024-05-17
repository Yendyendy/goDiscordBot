package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

type javData struct {
	Actresses       []actresses  `json:"actresses"`
	Categories      []categories `json:"categories"`
	Content_id      string       `json:"content_id"`
	Director        string       `json:"director"`
	Images          images       `json:"images"`
	Label           label        `json:"label"`
	Maker           maker        `json:"maker"`
	Release_date    string       `json:"release_date"`
	Runtime_minutes string       `json:"runtime_minutes"`
	Sample          sample       `json:"sample"`
	Series          series       `json:"series"`
	Title           string       `json:"title"`
}

type actresses struct {
	Name string `json:"name"`
}

type categories struct {
	Name string `json:"name"`
}

type images struct {
	Jacket_image jacket_image `json:"jacket_image"`
}

type jacket_image struct {
	Large  string `json:"large"`
	Large2 string `json:"large2"`
}

type label struct {
	Name string `json:"name"`
}

type maker struct {
	Name string `json:"name"`
}

type sample struct {
	High string `json:"high"`
}

type series struct {
	Name       string `json:"name"`
	Series_url string `json:"series_url"`
}

var (
	Token   = flag.String("t", "", "Bot authentication token")
	App     = flag.String("a", "", "Application ID")
	Channel = flag.String("c", "", "Server channer ID")
)

func main() {
	flag.Parse()
	if *App == "" {
		log.Fatal("err:: Application id is not set")
	}

	session, err := discordgo.New("Bot " + *Token)
	if err != nil {
		log.Fatal("err:: func NEW")
	}

	session.AddHandler(messageCreate)
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	ch := make(chan string)

	if m.Author.ID == s.State.User.ID {
		return
	}
	str := m.Content
	matched, _ := regexp.MatchString("[[a-zA-Z]{3,4}-[0-9]{3,5}]", m.Content)

	if matched {
		t := strings.Replace(str, "[", "", -1)
		t2 := strings.Replace(t, "]", "", -1)

		go getJavDataJson(t2, ch)

		msg := <-ch
		s.ChannelMessageSend(*Channel, msg)

	} else {
		s.ChannelMessageSend(*Channel, "No valid code")
	}

}

func getJavDataJson(code string, ch chan string) {
	c := colly.NewCollector()

	str := fmt.Sprintf("https://r18.dev/videos/vod/movies/detail/-/dvd_id=%s/json", code)

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", str)
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		str := string(r.Body)
		res := javData{}
		json.Unmarshal([]byte(str), &res)

		ch <- fmt.Sprintf("https://r18.dev/videos/vod/movies/detail/-/id=%s/", res.Content_id)
	})

	c.Visit(str)
}

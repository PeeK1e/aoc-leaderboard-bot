package main

import (
	"time"

	"github.com/PeeK1e/aoc-leaderboard-bot/aoc-bot/internal/config"
	"github.com/PeeK1e/aoc-leaderboard-bot/aoc-bot/internal/scrape"
	"github.com/PeeK1e/aoc-leaderboard-bot/aoc-bot/internal/webhook"
	"github.com/charmbracelet/log"
)

func main() {
	c := config.New()

	log.SetReportCaller(true)
	log.SetLevel(c.GetDebugLevel())

	log.Debug("config", "c", c)

	webhook.SetWebhookUrl(*c.WebHookUrl)

	t := time.Tick(900 * time.Second)

	if err := scrape.LoadFile(*c.DatabasePath); err != nil {
		log.Error("Can't read or load file", "err", err, "file", *c.DatabasePath)
	}

	for {
		if err := scrape.Scrape(c.GetUrl(), c.GetCookieToken()); err != nil {
			log.Error("Error scraping the Leader Board", "err", err)
		}

		webhook.Run()

		scrape.WriteToFile(*c.DatabasePath)
		<-t
	}
}

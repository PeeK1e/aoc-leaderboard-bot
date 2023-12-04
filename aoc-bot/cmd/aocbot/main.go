package main

import (
	"github.com/PeeK1e/aoc-leaderboard-bot/aoc-bot/internal/config"
	"github.com/PeeK1e/aoc-leaderboard-bot/aoc-bot/internal/database"
	"github.com/charmbracelet/log"
)

func main() {
	c := config.New()

	log.SetReportCaller(true)
	log.SetLevel(c.GetDebugLevel())

	log.Debug("config", "c", c)

	err := database.Open(*c.DatabasePath)
	if err != nil {
		log.Error("couldn't open db", "err", err)
	}

	defer database.Close()

}

package webhook

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PeeK1e/aoc-leaderboard-bot/aoc-bot/internal/scrape"
	"github.com/charmbracelet/log"
)

var webhookUrl string

type star struct {
	Day    string
	Name   string
	Ts     int64
	Amount string
}

func SetWebhookUrl(url string) {
	webhookUrl = url
}

func Run() {
	d := scrape.GetData()

	stamp := d.TimeStamp

	for _, v := range d.LeaderBoard.Members {
		stars := getMemberStars(v, stamp)
		createNotifications(stars)
	}
}

func getMemberStars(m scrape.Member, t int64) []star {
	stars := make([]star, 0)
	c := 0
	for day, v := range m.CompletionDayLevel {
		for starNo, s := range v {
			if s.GetStarTs > t {
				stars = append(stars, star{
					Day:    day,
					Name:   m.Name,
					Ts:     s.GetStarTs,
					Amount: starNo,
				})
			}
		}
		c++
	}
	return stars
}

func createNotifications(l []star) {
	for _, v := range l {
		log.Debugf("star:\n\tname:%s\n\tstars:%s\n\ttimestamp:%d", v.Name, v.Amount, v.Ts)
		s, err := createEmbed(v)
		if err != nil {
			log.Error("cannot create embed message", "err", err)
			continue
		}
		sendNotification(s)
	}
}

func sendNotification(s string) error {
	r, err := http.Post(webhookUrl, "application/json", bytes.NewBufferString(s))
	if err != nil {
		log.Error("can't post embed", "err", err)
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Error("webhook post error", "err", r.Status)
		return fmt.Errorf("webhook post error %s", r.Status)
	}

	return nil
}

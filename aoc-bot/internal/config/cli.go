package config

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/charmbracelet/log"
)

type Config struct {
	URL          *string
	CookieToken  *string
	DebugLevel   *string
	DatabasePath *string
	WebHookUrl   *string
}

func New() Config {
	c := Config{
		URL:          kingpin.Flag("url", "url of the leaderboard").Required().Short('u').Envar("AOC_URL").String(),
		CookieToken:  kingpin.Flag("cookie", "cookie needed for authentication").Short('c').Envar("AOC_COOKIE").Required().String(),
		DebugLevel:   kingpin.Flag("debug-level", "set the debug level you need").Short('l').Envar("AOC_LOG").Default("INFO").Enum("DEBUG", "INFO", "WARN", "ERROR"),
		DatabasePath: kingpin.Flag("db-path", "path of the sqlite3 db").Short('p').Envar("AOC_DB_PATH").Default("./db/db.sqlite").String(),
		WebHookUrl:   kingpin.Flag("webhook-url", "webhook to send the star info to").Required().Short('w').String(),
	}

	kingpin.Parse()

	return c
}

func (c Config) GetUrl() string {
	return *c.URL
}

func (c Config) GetCookieToken() string {
	return *c.CookieToken
}

func (c Config) GetDebugLevel() log.Level {
	if l, err := log.ParseLevel(*c.DebugLevel); err != nil {
		log.Error("error parsing level", "err", err)
		return log.DebugLevel
	} else {
		return l
	}
}

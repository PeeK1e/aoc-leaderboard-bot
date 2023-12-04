package webhook

import (
	"bytes"
	"text/template"
	"time"

	"github.com/charmbracelet/log"
)

func createEmbed(s star) (string, error) {
	const tmpl = `{
		"embeds": [{
			"title": "{{ .Name }} achieved new star(s)! ⭐",
			"description": "**They solved Day {{ .Day }}'s puzzle with {{ .Amount }} Stars!**",
			"timestamp": "{{ convert .Ts }}"
		}]
	}`

	t, err := template.New("embed").Funcs(template.FuncMap{
		"convert": func(ts int64) string {
			return time.Unix(ts, 0).Format(time.RFC3339)
		},
	}).Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, s); err != nil {
		return "", err
	}

	log.Debug("templated message", "message", buf.String())

	return buf.String(), nil
}

package scrape

import (
	"io"
	"net/http"
)

func Scrape(url string, cookie string) error {
	c := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	ck := http.Cookie{
		Name:  "session",
		Value: cookie,
	}

	req.AddCookie(&ck)

	r, err := c.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	d, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := saveData(&d); err != nil {
		return err
	}

	return nil
}

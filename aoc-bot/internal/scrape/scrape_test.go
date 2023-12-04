package scrape

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// the sample crawled webpage
var cURLsample string = `{"members":{"3398843":{"last_star_ts":1701602011,"global_score":0,"local_score":21,"stars":6,"id":3398843,"name":"dj95","completion_day_level":{"3":{"2":{"get_star_ts":1701602011,"star_index":558793},"1":{"get_star_ts":1701598804,"star_index":547762}},"2":{"2":{"get_star_ts":1701527277,"star_index":359807},"1":{"get_star_ts":1701526290,"star_index":356152}},"1":{"2":{"get_star_ts":1701448445,"star_index":112042},"1":{"star_index":83967,"get_star_ts":1701441526}}}},"3331156":{"stars":3,"id":3331156,"last_star_ts":1701675955,"global_score":0,"local_score":8,"completion_day_level":{"1":{"1":{"star_index":0,"get_star_ts":1701418546},"2":{"get_star_ts":1701537665,"star_index":397615}},"2":{"1":{"get_star_ts":1701675955,"star_index":777000}}},"name":"thled"},"2468679":{"stars":0,"id":2468679,"global_score":0,"last_star_ts":0,"local_score":0,"completion_day_level":{},"name":"Mawarii"},"1521470":{"local_score":13,"last_star_ts":1701522566,"global_score":0,"id":1521470,"stars":4,"name":"PeeK1e","completion_day_level":{"1":{"1":{"star_index":109289,"get_star_ts":1701447750},"2":{"get_star_ts":1701454443,"star_index":134799}},"2":{"1":{"get_star_ts":1701521091,"star_index":336952},"2":{"get_star_ts":1701522566,"star_index":342347}}}}},"event":"2023","owner_id":3398843}`

func TestScraper(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(cURLsample))
	}))
	defer server.Close()

	if err := Scrape(server.URL, ""); err != nil {
		t.Error(err)
	}
}

// this scrapes, saves, writes to a file and loads it again
// it should return the same as the input
func TestSaveAndLoad(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(cURLsample))
	}))
	defer server.Close()

	path := fmt.Sprintf("%s/store.json", t.TempDir())

	if err := Scrape(server.URL, ""); err != nil {
		t.Error(err)
	}

	err := WriteToFile(path)
	if err != nil {
		t.Error(err)
	}

	err = LoadFile(path)
	if err != nil {
		t.Error(err)
	}

	lb := new(LeaderBoard)
	if err := json.Unmarshal([]byte(cURLsample), lb); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(*lb, data.LeaderBoard) {
		t.Error("JSON are not equal")
	}
}

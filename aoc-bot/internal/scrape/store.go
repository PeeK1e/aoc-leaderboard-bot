package scrape

import (
	"encoding/json"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var data *StoredData = new(StoredData)

func GetData() StoredData {
	copy := *data
	return copy
}

func WriteToFile(path string) error {
	t := time.Now().Unix()

	data.TimeStamp = t

	log.Debug("Writing data to File", "path", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	log.Debug("saved file")
	return nil
}

func LoadFile(path string) error {
	log.Debug("loading data from file", "file", path)
	d, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(d, data)
	if err != nil {
		return err
	}

	log.Debug("loaded data")
	return nil
}

func saveData(b *[]byte) error {
	log.Debug("storing data")

	lb := &LeaderBoard{}
	err := json.Unmarshal(*b, lb)
	if err != nil {
		return err
	}

	data.LeaderBoard = *lb

	return nil
}

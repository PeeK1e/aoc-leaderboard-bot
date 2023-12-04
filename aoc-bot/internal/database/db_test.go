package database

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	path := fmt.Sprintf("%s/db.sqlite", t.TempDir())
	if err := Open(path); err != nil {
		t.Error(err)
	}
}

func TestExec(t *testing.T) {
	path := fmt.Sprintf("%s/db.sqlite", t.TempDir())
	if err := Open(path); err != nil {
		t.Error(err)
	}

	s := "INSERT INTO users (name, day, stars) VALUES('peek1e', '2', '1')"
	if _, err := RunQuery(s); err != nil {
		t.Error(err)
	}
}

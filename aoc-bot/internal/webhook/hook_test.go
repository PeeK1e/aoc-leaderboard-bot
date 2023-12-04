package webhook

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPostEmbed(t *testing.T) {
	// Create a dummy server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate method and header
		if r.Method != http.MethodPost {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected 'Content-Type' of 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}

		// Write response
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	SetWebhookUrl(server.URL)

	// Prepare data
	s := star{
		Day:    "2",
		Name:   "Alice",
		Ts:     time.Now().Unix(),
		Amount: "2",
	}

	embed, err := createEmbed(s)
	if err != nil {
		t.Fatal(err)
	}

	sendNotification(embed)
	if err != nil {
		t.Error(err)
	}
}

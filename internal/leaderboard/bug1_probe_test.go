package leaderboard

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProbe_EmptyTopReturnsArray(t *testing.T) {
	srv := httptest.NewServer(NewAPI().Handler())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/top?n=10")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status=%d want 200", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"entries":[]`) {
		t.Fatalf("empty top body=%s, want entries to be an empty JSON array", body)
	}
}

package rest

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestResponseBytesAndString(t *testing.T) {
  tmux := http.NewServeMux()
  server := httptest.NewServer(tmux)

  resp := Get(server.URL + "/user")

  if resp.StatusCode != http.StatusOK {
    t.Fatal("Status != OK (200)")
  }

  if string(resp.Bytes()) != resp.String() {
    t.Fatal("Bytes() and String() are not equal")
  }
}
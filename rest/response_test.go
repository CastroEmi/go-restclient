package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestDebug(t *testing.T) {

	resp := Get(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	if !strings.Contains(resp.Debug(), resp.String()) {
		t.Fatal("Debug() failed!")
	}

}

func TestGetFillUpJSON(t *testing.T) {

	var u []User

	resp := rb.Get("/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	err := resp.FillUp(&u)
	if err != nil {
		t.Fatal("Json fill up failed. Error: " + err.Error())
	}

	for _, v := range users {
		if v.Name == "Hernan" {
			return
		}
	}

	t.Fatal("Couldn't found Hernan")
}

func TestGetFillUpXML(t *testing.T) {

	var u []User

	var rbXML = RequestBuilder{
		BaseURL:     server.URL,
		ContentType: XML,
	}

	resp := rbXML.Get("/xml/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	err := resp.FillUp(&u)
	if err != nil {
		t.Fatal("Json fill up failed. Error: " + err.Error())
	}

	for _, v := range users {
		if v.Name == "Hernan" {
			return
		}
	}

	t.Fatal("Couldn't found Hernan")
}

package jiangjing

import (
	"encoding/json"
	"net/http"
	"testing"
)

const (
	address  = "http://localhost:3002"
	username = "elastic"
	password = "elastic823"
)

func TestGetHealth(t *testing.T) {
	client, err := NewClient(Config{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}

	resp, err := client.Health()
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, resp.StatusCode)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		t.Fatalf("Error parsing the response body: %s", err)
	}

	switch name := r["name"].(type) {
	case string:
		if len(name) == 0 {
			t.Fatal("Expect to have a field `name` in response.")
		}
	default:
		t.Fatal("Expect to have a field `name` in response.")
	}

	_, ok := r["version"]
	if !ok {
		t.Fatal("Expect to have a field `version` in response.")
	}
}
